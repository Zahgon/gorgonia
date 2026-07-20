package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"net/http"
	_ "net/http/pprof"

	G "gorgonia.org/gorgonia"
	"gorgonia.org/gorgonia/examples/mnist"
	"gorgonia.org/tensor"

	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

var (
	epochs     = flag.Int("epochs", 100, "Number of epochs to train for")
	dataset    = flag.String("dataset", "train", "Which dataset to train on? Valid options are \"train\" or \"test\"")
	dtype      = flag.String("dtype", "float64", "Which dtype to use")
	batchsize  = flag.Int("batchsize", 100, "Batch size")
	cpuprofile = flag.String("cpuprofile", "", "CPU profiling")
)

const loc = "../testdata/mnist/"

var dt tensor.Dtype

func parseDtype() { _ = "STUB: not implemented"; return }

type convnet struct {
	g                  *G.ExprGraph
	w0, w1, w2, w3, w4 *G.Node
	d0, d1, d2, d3     float64

	out *G.Node
}

func newConvNet(g *G.ExprGraph) *convnet { _ = "STUB: not implemented"; return nil }

func (m *convnet) learnables() G.Nodes { _ = "STUB: not implemented"; return *new(G.Nodes) }

func (m *convnet) fwd(x *G.Node) (err error) { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Parse()
	parseDtype()
	rand.Seed(1337)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	doneChan := make(chan bool, 1)

	var inputs, targets tensor.Tensor
	var err error

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	trainOn := *dataset
	if inputs, targets, err = mnist.Load(trainOn, loc, dt); err != nil {
		log.Fatal(err)
	}

	numExamples := inputs.Shape()[0]
	bs := *batchsize

	if err := inputs.Reshape(numExamples, 1, 28, 28); err != nil {
		log.Fatal(err)
	}
	g := G.NewGraph()
	x := G.NewTensor(g, dt, 4, G.WithShape(bs, 1, 28, 28), G.WithName("x"))
	y := G.NewMatrix(g, dt, G.WithShape(bs, 10), G.WithName("y"))
	m := newConvNet(g)
	if err = m.fwd(x); err != nil {
		log.Fatalf("%+v", err)
	}

	losses := G.Must(G.Log(G.Must(G.HadamardProd(m.out, y))))
	cost := G.Must(G.Mean(losses))
	cost = G.Must(G.Neg(cost))

	var costVal G.Value
	G.Read(cost, &costVal)

	if _, err = G.Grad(cost, m.learnables()...); err != nil {
		log.Fatal(err)
	}

	prog, locMap, _ := G.Compile(g)

	vm := G.NewTapeMachine(g, G.WithPrecompiled(prog, locMap), G.BindDualValues(m.learnables()...))
	solver := G.NewRMSPropSolver(G.WithBatchSize(float64(bs)))
	defer vm.Close()

	var profiling bool
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		profiling = true
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go cleanup(sigChan, doneChan, profiling)

	batches := numExamples / bs
	log.Printf("Batches %d", batches)
	bar := pb.New(batches)
	bar.SetRefreshRate(time.Second)
	bar.SetMaxWidth(80)

	for i := 0; i < *epochs; i++ {
		bar.Prefix(fmt.Sprintf("Epoch %d", i))
		bar.Set(0)
		bar.Start()
		for b := 0; b < batches; b++ {
			start := b * bs
			end := start + bs
			if start >= numExamples {
				break
			}
			if end > numExamples {
				end = numExamples
			}

			var xVal, yVal tensor.Tensor
			if xVal, err = inputs.Slice(G.S(start, end)); err != nil {
				log.Fatal("Unable to slice x")
			}

			if yVal, err = targets.Slice(G.S(start, end)); err != nil {
				log.Fatal("Unable to slice y")
			}
			if err = xVal.(*tensor.Dense).Reshape(bs, 1, 28, 28); err != nil {
				log.Fatalf("Unable to reshape %v", err)
			}

			G.Let(x, xVal)
			G.Let(y, yVal)
			if err = vm.RunAll(); err != nil {
				log.Fatalf("Failed at epoch  %d, batch %d. Error: %v", i, b, err)
			}
			if err = solver.Step(G.NodesToValueGrads(m.learnables())); err != nil {
				log.Fatalf("Failed to update nodes with gradients at epoch %d, batch %d. Error %v", i, b, err)
			}
			vm.Reset()
			bar.Increment()
		}
		log.Printf("Epoch %d | cost %v", i, costVal)

	}
}

func cleanup(sigChan chan os.Signal, doneChan chan bool, profiling bool) {
	_ = "STUB: not implemented"
	return
}

func handlePprof(sigChan chan os.Signal, doneChan chan bool) { _ = "STUB: not implemented"; return }
