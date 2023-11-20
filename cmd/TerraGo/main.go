package main

import (
	"github.com/Niviann/TerraGo/internal/noise"
	"github.com/Niviann/TerraGo/internal/output"
)

func main() {
	var seed uint64 = 12345678901234566666
	var chunk_size int = 2048
	var octaves int = 4
	output.CreateImage(noise.MapGen(seed, chunk_size, octaves), "noise")
	output.ColorImage(noise.MapGen(seed, chunk_size, octaves), "color")
}
