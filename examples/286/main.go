package main

import (
	"flag"
	"log"
	"math/rand"

	_ "net/http/pprof"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var (
	epochs     = flag.Int("epochs", 10, "Number of epochs to train for")
	dataset    = flag.String("dataset", "train", "Which dataset to train on? Valid options are \"train\" or \"test\"")
	dtype      = flag.String("dtype", "float64", "Which dtype to use")
	batchsize  = flag.Int("batchsize", 10, "Batch size")
	cpuprofile = flag.String("cpuprofile", "", "CPU profiling")
)

const loc = "./mnist/"

var dt tensor.Dtype

func parseDtype() { _ = "STUB: not implemented"; return }

type nn struct {
	g      *gorgonia.ExprGraph
	w0, w1 *gorgonia.Node

	out     *gorgonia.Node
	predVal gorgonia.Value
}

type sli struct {
	start, end int
}

func (s sli) Start() int { _ = "STUB: not implemented"; return 0 }
func (s sli) End() int   { _ = "STUB: not implemented"; return 0 }
func (s sli) Step() int  { _ = "STUB: not implemented"; return 0 }

func newNN(g *gorgonia.ExprGraph) *nn { _ = "STUB: not implemented"; return nil }

func (m *nn) learnables() gorgonia.Nodes { _ = "STUB: not implemented"; return *new(gorgonia.Nodes) }

func (m *nn) fwd(x *gorgonia.Node) (err error) { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Parse()
	parseDtype()
	rand.Seed(7945)

	var err error

	bs := *batchsize
	g := gorgonia.NewGraph()
	x := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(bs, 784), gorgonia.WithName("x"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	y := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(bs, 10), gorgonia.WithName("y"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))

	m := newNN(g)
	if err = m.fwd(x); err != nil {
		log.Fatalf("%+v", err)
	}

	losses, err := gorgonia.HadamardProd(m.out, y)
	if err != nil {
		log.Fatal(err)
	}
	cost := gorgonia.Must(gorgonia.Mean(losses))
	cost = gorgonia.Must(gorgonia.Neg(cost))

	var costVal gorgonia.Value
	gorgonia.Read(cost, &costVal)

	if _, err = gorgonia.Grad(cost, m.learnables()...); err != nil {
		log.Fatal(err)
	}

	vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(m.learnables()...))
	solver := gorgonia.NewRMSPropSolver(gorgonia.WithBatchSize(float64(bs)))
	defer vm.Close()

	if err = vm.RunAll(); err != nil {
		log.Fatalf("Failed %v", err)
	}

	solver.Step(gorgonia.NodesToValueGrads(m.learnables()))
	vm.Reset()
}
