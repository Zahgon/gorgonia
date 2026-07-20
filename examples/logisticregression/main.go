package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

const (
	N = 26733

	feats = 10

	trainIter = 500
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write mem profile to file")
var static = flag.Bool("static", false, "Use static test file")
var wT tensor.Tensor
var yT tensor.Tensor
var xT tensor.Tensor

var Float = tensor.Float64

func init() {
	xBacking := tensor.Random(Float, N*feats)
	wBacking := tensor.Random(Float, feats)
	var yBacking interface{}
	switch Float {
	case tensor.Float64:
		backing := make([]float64, N)
		for i := range backing {
			backing[i] = float64(rand.Intn(2))
		}
		yBacking = backing
	case tensor.Float32:
		backing := make([]float32, N)
		for i := range backing {
			backing[i] = float32(rand.Intn(2))
		}
		yBacking = backing
	}

	xT = tensor.New(tensor.WithBacking(xBacking), tensor.WithShape(N, feats))
	yT = tensor.New(tensor.WithBacking(yBacking), tensor.WithShape(N))
	wT = tensor.New(tensor.WithBacking(wBacking), tensor.WithShape(feats))
}

func main() {
	flag.Parse()
	rand.Seed(1337)
	log.SetFlags(0)

	if *static {
		Float = tensor.Float64
		wBacking, xBacking, yBacking := loadStatic()
		xT = tensor.New(tensor.WithBacking(xBacking), tensor.WithShape(N, feats))
		yT = tensor.New(tensor.WithBacking(yBacking), tensor.WithShape(N))
		wT = tensor.New(tensor.WithBacking(wBacking), tensor.WithShape(feats))
	}

	g := G.NewGraph()
	x := G.NewMatrix(g, Float, G.WithName("x"), G.WithShape(N, feats))
	y := G.NewVector(g, Float, G.WithName("y"), G.WithShape(N))

	w := G.NewVector(g, Float, G.WithName("w"), G.WithShape(feats))
	b := G.NewScalar(g, Float, G.WithName("bias"))

	one := G.NewConstant(1.0)

	xwmb := G.Must(G.Add(G.Must(G.Mul(x, w)), b))

	prob := G.Must(G.Sigmoid(xwmb))
	G.WithName("prob")(prob)

	pred := G.Must(G.Gt(prob, G.NewConstant(0.5), false))
	G.WithName("pred")(pred)

	var predicted G.Value
	readNode := G.Read(pred, &predicted)

	logProb := G.Must(G.Log(prob))
	fstTerm := G.Must(G.HadamardProd(G.Must(G.Neg(y)), logProb))
	oneMinusY := G.Must(G.Sub(one, y))
	logOneMinusProb := G.Must(G.Log(G.Must(G.Sub(one, prob))))
	sndTerm := G.Must(G.HadamardProd(oneMinusY, logOneMinusProb))

	crossEntropy := G.Must(G.Sub(fstTerm, sndTerm))
	G.WithName("crossEntropy")(crossEntropy)
	loss := G.Must(G.Mean(crossEntropy))
	G.WithName("loss")(loss)

	weightSq := G.Must(G.Square(w))
	sumSq := G.Must(G.Sum(weightSq))
	l2reg := G.NewConstant(0.01, G.WithName("l2reg"))
	regTerm := G.Must(G.Mul(l2reg, sumSq))

	cost := G.Must(G.Add(loss, regTerm))
	G.WithName("cost")(cost)

	grads, err := G.Grad(cost, w, b)
	handleError(err)

	G.WithName("dcost/dw")(grads[0])

	G.WithName("dcost/db")(grads[1])

	learnRate := G.NewConstant(0.1)
	gwlr := G.Must(G.Mul(learnRate, grads[0]))
	wUpd := G.Must(G.Sub(w, gwlr))
	gblr := G.Must(G.Mul(learnRate, grads[1]))
	bUpd := G.Must(G.Sub(b, gblr))

	prog, locMap, err := G.CompileFunction(g, G.Nodes{x, y}, G.Nodes{wUpd, bUpd, readNode})
	handleError(err)
	fmt.Printf("%v", prog)

	machine := G.NewTapeMachine(g, G.WithPrecompiled(prog, locMap))
	defer machine.Close()

	machine.Let(w, wT)
	machine.Let(b, 0.0)

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	for i := 0; i < trainIter; i++ {

		machine.Reset()

		machine.Let(x, xT)
		machine.Let(y, yT)
		handleError(machine.RunAll())

		machine.Set(w, wUpd)
		machine.Set(b, bUpd)

		accuracy := accuracy(y.Value(), predicted)
		fmt.Printf("Interation #%v, Training accuracy: %#v\n", i, accuracy)

	}
	fmt.Printf("Time taken: %v\n", time.Since(start))
	fmt.Printf("Final Model: \nw: %3.3s\nb: %+3.3s\n", w.Value(), b.Value())

	fmt.Printf("START\n")

	prog, locMap, err = G.CompileFunction(g, G.Nodes{x}, G.Nodes{pred})
	handleError(err)
	machine = G.NewTapeMachine(g, G.WithPrecompiled(prog, locMap))
	defer machine.Close()

	machine.Let(w, wT)
	machine.Let(b, 0.0)
	machine.Let(x, xT)
	handleError(machine.RunAll())
	handleError(err)

}
func accuracy(target, predicted G.Value) float64 { _ = "STUB: not implemented"; return 0 }

func handleError(err error) { _ = "STUB: not implemented"; return }
