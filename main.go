package main

var chunk_size int = 1024

type Vector2 struct {
	x, y float64
}

func main() {
	create_image(mapGen(21221, 23), "img")
}
