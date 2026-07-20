package main

type LayerConfig struct {
	Inputs, Outputs int
	BatchSize       int
}

type DeepConfig struct {
	LayerConfig

	Size              int
	Layers            int
	HiddenLayersSizes []int
}
