package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func create_image(slice [][]float64, name string) {
	// create image from slice
	img := image.NewGray(image.Rect(0, 0, len(slice), len(slice)))
	for y := 0; y < len(slice); y++ {
		for x := 0; x < len(slice); x++ {
			grayValue := uint8((slice[x][y] + 1) / 2 * 255)
			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	outputFile, err := os.Create(name + ".png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	png.Encode(outputFile, img)
}
