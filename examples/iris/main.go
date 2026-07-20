package main

import (
	"fmt"
	"log"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/mat"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	g := gorgonia.NewGraph()
	x, y := getXYMat()
	xT := tensor.FromMat64(mat.DenseCopyOf(x))
	yT := tensor.FromMat64(mat.DenseCopyOf(y))

	s := yT.Shape()
	yT.Reshape(s[0])

	X := gorgonia.NodeFromAny(g, xT, gorgonia.WithName("x"))
	Y := gorgonia.NodeFromAny(g, yT, gorgonia.WithName("y"))
	theta := gorgonia.NewVector(
		g,
		gorgonia.Float64,
		gorgonia.WithName("theta"),
		gorgonia.WithShape(xT.Shape()[1]),
		gorgonia.WithInit(gorgonia.Gaussian(0, 1)))

	pred := must(gorgonia.Mul(X, theta))

	var predicted gorgonia.Value
	gorgonia.Read(pred, &predicted)

	squaredError := must(gorgonia.Square(must(gorgonia.Sub(pred, Y))))
	cost := must(gorgonia.Mean(squaredError))

	if _, err := gorgonia.Grad(cost, theta); err != nil {
		log.Fatalf("Failed to backpropagate: %v", err)
	}

	machine := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(theta))
	defer machine.Close()

	model := []gorgonia.ValueGrad{theta}
	solver := gorgonia.NewVanillaSolver(gorgonia.WithLearnRate(0.001))

	fa := mat.Formatted(getThetaNormal(x, y), mat.Prefix("   "), mat.Squeeze())

	fmt.Printf("ϴ: %v\n", fa)
	iter := 10000
	var err error
	for i := 0; i < iter; i++ {
		if err = machine.RunAll(); err != nil {
			fmt.Printf("Error during iteration: %v: %v\n", i, err)
			break
		}

		if err = solver.Step(model); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("theta: %2.2f  Iter: %v Cost: %2.3f Accuracy: %2.2f \r",
			theta.Value(),
			i,
			cost.Value(),
			accuracy(predicted.Data().([]float64), Y.Value().Data().([]float64)))

		machine.Reset()
	}
	fmt.Println("")
	err = save(theta.Value())
	if err != nil {
		log.Fatal(err)
	}

}

func accuracy(prediction, y []float64) float64 { _ = "STUB: not implemented"; return 0 }

func getXYMat() (*matrix, *matrix) { _ = "STUB: not implemented"; return nil, nil }

func getThetaNormal(x, y *matrix) *mat.Dense { _ = "STUB: not implemented"; return nil }

type matrix struct {
	dataframe.DataFrame
}

func (m matrix) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (m matrix) T() mat.Matrix { _ = "STUB: not implemented"; return *new(mat.Matrix) }

func must(n *gorgonia.Node, err error) *gorgonia.Node { _ = "STUB: not implemented"; return nil }

func one(size int) []float64 { _ = "STUB: not implemented"; return nil }

func save(value gorgonia.Value) error { _ = "STUB: not implemented"; return nil }
