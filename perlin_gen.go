package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

// My implementation of perlin noise generator

var seed int64 = 1234567890
var chunk_size int = 32
var map_size int = 10

func interpolate(a0, a1, w float64) float64 {
	return (a1-a0)*w + a0
}

type Vector2 struct {
	x, y float64
}

type Vector3 struct {
	X, Y, Z float64
}

func dot(g Vector2, x, y float64) float64 {
	var vector_length float64 = math.Sqrt(float64(math.Pow(float64(x+1), 2) + math.Pow(float64(y+1), 2)))
	return g.x*(x+1)/vector_length + g.y*(y+1)/vector_length
}

func randomRotatedVector2() Vector2 {
	var vector Vector2
	rng := rand.New(rand.NewSource(seed))
	vector.x, vector.y = math.Sincos(rng.Float64() * math.Pi * 2)
	return vector
}

func randomGradient() {
	//create 2 dimensional slice that will store chunk data
	var chunk [][]float64
	var vector Vector2 = randomRotatedVector2()
	for i := 0; i < chunk_size; i++ {
		var row []float64
		for j := 0; j < chunk_size; j++ {
			row = append(row, dot(vector, float64(j), float64(i)))
		}
		chunk = append(chunk, row)
	}

	img := image.NewGray(image.Rect(0, 0, chunk_size, chunk_size))
	for y := 0; y < chunk_size; y++ {
		for x := 0; x < chunk_size; x++ {
			// Assuming float values are in the range [0, 1]
			grayValue := uint8((chunk[x][y] + 1) / 2 * 255)
			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	outputFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, img)

}
