package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"maxGroupOfOneColor/matrix"
)

const (
	defaultMatrixRows    = 10
	defaultMatrixCols    = 10
	defaultCountOfColors = 3
)

func main() {

	rand.Seed(time.Now().UnixNano())

	rows := flag.Int("r", defaultMatrixRows, "count of matrix rows")
	cols := flag.Int("c", defaultMatrixCols, "count of matrix cols")
	colors := flag.Int("cc", defaultCountOfColors, "count of colors")

	flag.Parse()

	if *rows <= 0 || *cols <= 0 || *colors <= 0 {
		log.Fatalf("Invalid params. Need rows > 0, cols > 0, color > 0")
	}

	mat := make([][]int, *rows)
	for row, _ := range mat {
		mat[row] = make([]int, *cols)
	}

	for row, _ := range mat {
		for col, _ := range mat[row] {
			mat[row][col] = rand.Intn(*colors)
		}
	}

	coordsOfMaxColorGroup, color := matrix.GetCoordsOfMaxColorGroup(mat)
	mapCoordsOfMaxColorGroup := make(map[int]struct{}, len(coordsOfMaxColorGroup))
	for _, coord := range coordsOfMaxColorGroup {
		mapCoordsOfMaxColorGroup[(coord.X-1)*len(mat[0])+coord.Y] = struct{}{}
	}

	GOOSIsLinux := runtime.GOOS == "linux"

	fmt.Printf("matrix:\n")
	coordForCheck := matrix.Coord{}
	for row, _ := range mat {
		fmt.Print("        ")
		for col, _ := range mat[row] {
			coordForCheck.X = row
			coordForCheck.Y = col
			if _, ok := mapCoordsOfMaxColorGroup[(row-1)*len(mat[0])+col]; ok && GOOSIsLinux {
				fmt.Printf("\033[01;38;05;196m%d\033[0m ", mat[row][col])
			} else {
				fmt.Printf("%d ", mat[row][col])
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	fmt.Printf("color: %d\ncount of fields: %d\ncoords: %v\n", color, len(coordsOfMaxColorGroup), coordsOfMaxColorGroup)
}
