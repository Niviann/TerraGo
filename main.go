package main

var chunk_size int = 2048
var map_seed uint64 = 10101010101010101010

type Vector2 struct {
	x, y float64
}

func main() {
	create_image(mapGen(), "img")

	color_image(mapGen())
}
