package main

import "C"

import (
	"flag"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"

	"gonum.org/v1/gonum/blas/gonum"
	T "gorgonia.org/gorgonia"
	"gorgonia.org/gorgonia/examples/mnist"
	"gorgonia.org/tensor"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var dataset = flag.String("dataset", "dev", "which data set to train on? Valid options: \"train\" or \"dev\"")
var bs = flag.Int("bs", 1, "training batch size (doesn't affect sgd batch size)")
var ptEpoch = flag.Int("pt", 16, "pretraining epoch")
var ftEpoch = flag.Int("ft", 40, "finetuning epoch")
var viz = flag.Int("viz", 0, "Visualize which layer?")
var save = flag.String("save", "", "Save file as")
var verbose = flag.Bool("v", false, "Verbose?")

var trainingWriter io.Writer
var trainingLog *log.Logger

var dt tensor.Dtype = tensor.Float64

func init() {
	var err error
	if trainingWriter, err = os.OpenFile("training.viz", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}

	trainingLog = log.New(trainingWriter, "", log.Ltime|log.Lmicroseconds)
}

func predictBatch(logprobs tensor.Tensor, batchSize int) (guesses []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func makeTargets(targets tensor.Tensor) []int { _ = "STUB: not implemented"; return nil }

func verboseLog(format string, attrs ...interface{}) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()
	rand.Seed(1337)

	loc := "../testdata/mnist/"

	trainOn := *dataset
	inputs, targets, err := mnist.Load(trainOn, loc, dt)
	if err != nil {
		log.Fatal(err)
	}

	size := inputs.Shape()[0]
	inputSize := 784
	outputSize := 10
	hiddenSizes := []int{1000, 1000, 1000}
	layers := len(hiddenSizes)
	corruptions := []float64{0.1, 0.2, 0.3}
	batchSize := *bs
	pretrainEpoch := *ptEpoch
	finetuneEpoch := *ftEpoch

	deets := `Stacked Denoising AutoEncoder
==============================
	Train on: %v
	Training Size: %v
	Hidden Sizes: %v
	Corruptions: %v
	Batch Size: %v
	Pretraining Epoch: %v
	Finetuning Epoch: %v
	BLAS used: %v
`
	fmt.Printf(deets, trainOn, size, hiddenSizes, corruptions, batchSize, pretrainEpoch, finetuneEpoch, T.WhichBLAS())

	g := T.NewGraph()
	sda := NewStackedDA(g, batchSize, size, inputSize, outputSize, layers, hiddenSizes, corruptions)

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	verboseLog("Pretraining...")
	for i := 0; i < pretrainEpoch; i++ {
		if err = sda.Pretrain(inputs, i); err != nil {
			ioutil.WriteFile("fullGraph_err.dot", []byte(g.ToDot()), 0644)
			log.Fatalf("i: %d err :%v", i, err)
		}
	}

	verboseLog("Starting to finetune now")

	T.Use(gonum.Implementation{})
	ys := makeTargets(targets)
	for i := 0; i < finetuneEpoch; i++ {
		if err = sda.Finetune(inputs, ys, i); err != nil {
			log.Fatal(err)
		}
	}

	if *save != "" {
		if err = sda.Save(*save); err != nil {
			log.Fatal(err)
		}
	}

	visualizeLayer := *viz
	verboseLog("Visualizing %dth layer", visualizeLayer)
	finalWeights := sda.autoencoders[visualizeLayer].w.Value().(tensor.Tensor).Clone().(tensor.Tensor)
	finalWeights.T()
	finalWeights.Transpose()
	for i := 0; i < finalWeights.Shape()[0]; i++ {
		rowT, _ := finalWeights.Slice(T.S(i))
		row := rowT.Data().([]float64)
		img := visualizeRow(row)

		f, _ := os.OpenFile(fmt.Sprintf("images/%d.jpg", i), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		f.Close()
	}

	verboseLog("pred")
	testX, testY, err := mnist.Load("test", loc, dt)
	if err != nil {
		log.Fatal(err)
	}
	var one, correct, lp tensor.Tensor
	if one, err = testX.Slice(T.S(0, batchSize)); err != nil {
		log.Fatal(err)
	}

	if correct, err = testY.Slice(T.S(0, batchSize)); err != nil {
		log.Fatal(err)
	}

	correctYs := makeTargets(correct)

	var predictions []int
	if lp, err = sda.Forwards(one); err != nil {
		log.Fatal(err)
	}

	if predictions, err = predictBatch(lp, batchSize); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Correct: \n%+v. \nPredicted: %v. \nLogprobs: \n%+#3.3s", correctYs, predictions, lp)
}
