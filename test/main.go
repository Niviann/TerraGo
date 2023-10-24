package main

import "math"

type Vector2 struct {
	X, Y float64
}

// random vector2 gen based on coordinates
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
	v := Vector2{X: math.Cos(random), Y: math.Sin(random)}
	return v
}

func main() {
	// Example usage:
	ix, iy := -123, 456
	result := randomVec2(ix, iy)
	println("X:", result.X)
	println("Y:", result.Y)
	println(math.MaxUint32)
}
