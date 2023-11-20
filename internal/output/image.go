package output

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func CreateImage(slice [][]float64, name string) {
	// create image from slice
	img := image.NewGray(image.Rect(0, 0, len(slice), len(slice)))
	for y := 0; y < len(slice); y++ {
		for x := 0; x < len(slice); x++ {
			grayValue := uint8((slice[x][y] + 1) / 2 * 255)
			img.SetGray(x, y, color.Gray{Y: grayValue})
		}
	}

	outputDir := "../../output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic("Error creating output directory: " + err.Error())
	}

	outputPath := filepath.Join(outputDir, name+".png")

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic("Error creating output file: " + err.Error())
	}
	defer outputFile.Close()

	// Encode the image as PNG and write it to the file
	if err := png.Encode(outputFile, img); err != nil {
		panic("Error encoding PNG: " + err.Error())
	}
}

func ColorImage(slice [][]float64, name string) {
	pixels := make([][]color.RGBA, len(slice))
	for y := 0; y < len(slice); y++ {
		pixels[y] = make([]color.RGBA, len(slice))
		for x := 0; x < len(slice); x++ {
			if slice[y][x] > .6 {
				pixels[y][x] = color.RGBA{255, 255, 255, 255}
			} else if slice[y][x] > .5 {
				pixels[y][x] = color.RGBA{102, 102, 153, 255}
			} else if slice[y][x] > .3 {
				pixels[y][x] = color.RGBA{0, 153, 51, 255}
			} else if slice[y][x] > .2 {
				pixels[y][x] = color.RGBA{128, 0, 0, 255}
			} else if slice[y][x] > .1 {
				pixels[y][x] = color.RGBA{255, 204, 102, 255}
			} else if slice[y][x] < -.25 {
				pixels[y][x] = color.RGBA{0, 0, 200, 255}
			} else {
				pixels[y][x] = color.RGBA{0, 0, 255, 255}
			}
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, len(slice), len(slice)))

	// Copy pixel values to the image
	for y := 0; y < len(slice); y++ {
		for x := 0; x < len(slice); x++ {
			img.Set(x, y, pixels[y][x])
		}
	}

	outputDir := "../../output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		panic("Error creating output directory: " + err.Error())
	}

	outputPath := filepath.Join(outputDir, name+".png")

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic("Error creating output file: " + err.Error())
	}
	defer outputFile.Close()

	// Encode the image as PNG and write it to the file
	if err := png.Encode(outputFile, img); err != nil {
		panic("Error encoding PNG: " + err.Error())
	}
}
