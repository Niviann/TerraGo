package main

import (
	"math"
)

// My implementation of perlin noise generator
var chunk_size int = 1024

// calculate dot product of influeance vector and offset vector
func dot(g Vector2, x, y int) float64 {
	var vector_length float64 = math.Sqrt(float64(math.Pow(float64(x+1), 2) + math.Pow(float64(y+1), 2)))
	return g.x*(float64(x)+1)/vector_length + g.y*(float64(y)+1)/vector_length
}

// create influence vector
func randomRotatedVector2() Vector2 {
	var vector Vector2
	vector.x, vector.y = math.Sincos(r.Float64() * math.Pi * 2)
	return vector
}

// bicubic interpolation
func interpolate(a0 float64, a1 float64, w float64) float64 {
	return ((a1-a0)*(3.0-w*2.0)*w*w + a0)
}

// func randomGradient(seed_val int) [][]float64 {
// 	// create 2 dimensional slice that will store dot products of random influeance vector and offset vectors
// 	var chunk [][]float64
// 	var vector Vector2 = randomRotatedVector2(seed_val) //create influeance vector by given seed

// 	for i := 0; i < chunk_size; i++ {
// 		var row []float64
// 		for j := 0; j < chunk_size; j++ {
// 			row = append(row, dot(vector, j, i))
// 		}
// 		chunk = append(chunk, row)
// 	}
// 	return chunk
// }

// generate chunk of perlin noise
func perlin(ch *chunk) {

	for i := 0; i < chunk_size; i++ {
		var row []float64
		for j := 0; j < chunk_size; j++ {

			sx := float64(j) / float64(chunk_size)
			sy := float64(i) / float64(chunk_size)

			n0 := dot(ch.vec1, j, i)
			n1 := dot(ch.vec2, j, i)
			ix0 := interpolate(n0, n1, sx)

			n0 = dot(ch.vec3, j, i)
			n1 = dot(ch.vec4, j, i)
			ix1 := interpolate(n0, n1, sx)

			value := interpolate(ix0, ix1, sy)

			row = append(row, value)
		}
		ch.perlin = append(ch.perlin, row)

	}
}

// 	for i := 0; i < len(chunk); i++ {
// 		for j := 0; j < len(chunk[i]); j++ {
// 			fmt.Printf("%f\t", chunk[i][j])
// 		}
// 		fmt.Println() // Move to the next row
// 	}
// 	return chunk
// }
