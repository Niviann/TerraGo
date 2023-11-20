package output

import (
	"fmt"
	"os"
)

func create_obj_file(slice [][]float64) {
	file, err := os.Create("3dMap.obj") //create a new file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for y := 0; y < len(slice)-1; y++ {
		for x := 0; x < len(slice[y])-1; x++ {

		}
	}
}
