package main

import (
	"math"
)

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
	const w = 32    // bit number
	const s = w / 2 // rotation width
	// ix += math.MaxUint32 / 2
	// iy += math.MaxUint32 / 2
	a, b := uint32(ix), uint32(iy)
	a *= 3284157443
	b ^= (a << s) | (a >> (w - s))
	b *= 1911520717
	a ^= (b << s) | (b >> (w - s))
	a *= 2048419325
	random := float64(a) * (math.Pi / float64(^uint32(0)>>1)) // in [0, 2*Pi]
	return Vector2{x: math.Cos(random), y: math.Sin(random)}
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

	value := interpolate(ix0, ix1, sy)
	return value
}

func mapGen(x, y float64) [][]float64 {
	var chunk [][]float64
	var GRID_SIZE float64 = 100
	for i := 0; i < chunk_size; i++ {
		var row []float64
		for j := 0; j < chunk_size; j++ {
			var freq float64 = 1
			var amp float64 = 1
			var val float64 = 0
			for k := 0; k < 10; k++ {
				val += perlin(float64(j)*freq/GRID_SIZE, float64(i)*freq/GRID_SIZE) * amp
				freq *= 2
				amp /= 2
			}
			row = append(row, val)

		}
		chunk = append(chunk, row)
	}

	return chunk
}
