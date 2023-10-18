package main

import (
	"math"
	"math/rand"
)

// My implementation of perlin noise generator

var seed int64 = 1234567890
var chunk_size int = 32
var map_size int = 10

type Vector2 struct {
	x, y float64
}

// calculate dot product of influeance vector and offset vector
func dot(g Vector2, x, y float64) float64 {
	var vector_length float64 = math.Sqrt(float64(math.Pow(float64(x+1), 2) + math.Pow(float64(y+1), 2)))
	return g.x*(x+1)/vector_length + g.y*(y+1)/vector_length
}

// used for generating influeance vectors
func randomRotatedVector2(seed_val int64) Vector2 {
	var vector Vector2
	rng := rand.New(rand.NewSource(seed_val))
	vector.x, vector.y = math.Sincos(rng.Float64() * math.Pi * 2)
	return vector
}

func randomGradient() [][]float64 {
	// create 2 dimensional slice that will store dot products of random influeance vector and offset vectors
	var chunk [][]float64
	var vector Vector2 = randomRotatedVector2(seed) //create influeance vector by given seed

	for i := 0; i < chunk_size; i++ {
		var row []float64
		for j := 0; j < chunk_size; j++ {
			row = append(row, dot(vector, float64(j), float64(i)))
		}
		chunk = append(chunk, row)
	}
	return chunk
}
