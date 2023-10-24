package main

import (
	"math"
)

<<<<<<< Updated upstream
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
=======
func dot(ix, iy int, x, y float64) float64 {
	vec := randomVec2(ix, iy)
	dx := x - float64(ix)
	dy := y - float64(iy)
	return dx*vec.x + dy*vec.y
}

// bicubic interpolation
func interpolate(a0 float64, a1 float64, w float64) float64 {
	return ((a1-a0)*(3.0-w*2.0)*w*w + a0)
}

func randomVec2(ix, iy int) Vector2 {
	// No precomputed gradients mean this works for any number of grid coordinates
	const w = 32
	const s = w / 2 // rotation width
	ix += math.MaxUint32 / 2
	iy += math.MaxUint32 / 2
	a := uint32(ix)
	b := uint32(iy)
	a *= 3284157443
	b ^= (a << s) | (a >> (w - s))
	b *= 1911520717
	a ^= (b << s) | (b >> (w - s))
	a *= 2048419325
	random := float64(a) * (math.Pi / float64(^uint32(0)>>1)) // in [0, 2*Pi]
	v := Vector2{x: math.Cos(random), y: math.Sin(random)}
	return v
}

func perlin(x, y float64) float64 {
	x0 := int(math.Floor(x))
	x1 := x0 + 1
	y0 := int(math.Floor(y))
	y1 := y0 + 1

	sx := x - float64(x0)
	sy := y - float64(y0)

	n0 := dot(x0, y0, x, y)
	n1 := dot(x1, y0, x, y)
	ix0 := interpolate(n0, n1, sx)

	n0 = dot(x0, y1, x, y)
	n1 = dot(x1, y1, x, y)
	ix1 := interpolate(n0, n1, sx)

	return interpolate(ix0, ix1, sy)
}

// func randomGradient() [][]float64 {
// 	// create 2 dimensional slice that will store dot products of random influeance vector and offset vectors
// 	var chunk [][]float64
// 	var vector Vector2 = randomRotatedVector2(seed) //create influeance vector by given seed
>>>>>>> Stashed changes

// 	for i := 0; i < chunk_size; i++ {
// 		var row []float64
// 		for j := 0; j < chunk_size; j++ {
<<<<<<< Updated upstream
// 			row = append(row, dot(vector, j, i))
=======
// 			row = append(row, dot(vector, float64(j), float64(i)))
>>>>>>> Stashed changes
// 		}
// 		chunk = append(chunk, row)
// 	}
// 	return chunk
// }

<<<<<<< Updated upstream
// generate chunk of perlin noise
func perlin(ch *chunk) {
=======
func mapGen(x, y float64) [][]float64 {
	var chunk [][]float64
>>>>>>> Stashed changes

	for i := 0; i < chunk_size; i++ {
		var row []float64
		for j := 0; j < chunk_size; j++ {
<<<<<<< Updated upstream

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
=======
			row = append(row, (perlin(float64(j)/float64(chunk_size)+x, float64(i)/float64(chunk_size))+1)/2+y)
>>>>>>> Stashed changes
		}
		ch.perlin = append(ch.perlin, row)

	}
<<<<<<< Updated upstream
=======

	return chunk
>>>>>>> Stashed changes
}

// 	for i := 0; i < len(chunk); i++ {
// 		for j := 0; j < len(chunk[i]); j++ {
// 			fmt.Printf("%f\t", chunk[i][j])
// 		}
// 		fmt.Println() // Move to the next row
// 	}
// 	return chunk
// }
