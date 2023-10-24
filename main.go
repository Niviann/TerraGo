package main

import "math/rand"

var seed int64 = 4123456275344
var r *rand.Rand = rand.New(rand.NewSource(seed))

type Vector2 struct {
	x, y float64
}

type pair struct {
	x, y int
}

type chunk struct {
	x, y                   int
	vec1, vec2, vec3, vec4 Vector2
	perlin                 [][]float64
}

func main() {
	chunks := make(map[pair]chunk)

	newChunk := chunk{
		x:    0,
		y:    0,
		vec1: randomRotatedVector2(),
		vec2: randomRotatedVector2(),
		vec3: randomRotatedVector2(),
		vec4: randomRotatedVector2(),
	}

	perlin(&newChunk)

	chunks[pair{0, 0}] = newChunk

	create_image(newChunk.perlin, "new")
}
