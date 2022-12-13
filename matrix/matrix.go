package matrix

type Coord struct {
	X int
	Y int
}

// GetCoordsOfMaxColorGroup returns the coordinates of the largest group of colors.
// If there is more than one largest group, then the one whose field is encountered first
// when traversing the matrix from top left to right down is returned.
func GetCoordsOfMaxColorGroup(matrix [][]int) (coordsOfMaxColorGroup []Coord, color int) {
	if len(matrix) == 0 {
		return nil, color
	}

	rows, cols := len(matrix), len(matrix[0])

	checkedCoords := make(map[int]struct{}, len(matrix)*len(matrix[0]))
	coordForCheck := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			coordForCheck = (row-1)*cols + col
			if _, ok := checkedCoords[coordForCheck]; !ok {
				currCoords := make([]Coord, 0)
				fillCoordsOfColorGroup(matrix, row, col, matrix[row][col], checkedCoords, &currCoords)
				if len(currCoords) > len(coordsOfMaxColorGroup) {
					color = matrix[row][col]
					coordsOfMaxColorGroup = currCoords
				}
			}
		}
	}

	return coordsOfMaxColorGroup, color
}

func fillCoordsOfColorGroup(matrix [][]int, row, col, color int, checkedCoords map[int]struct{}, coords *[]Coord) {
	if len(matrix) == 0 {
		return
	}
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] != color {
		return
	}
	coordForCheck := (row-1)*len(matrix[0]) + col
	if _, ok := checkedCoords[coordForCheck]; ok {
		return
	}
	checkedCoords[coordForCheck] = struct{}{}

	*coords = append(*coords, Coord{row, col})

	fillCoordsOfColorGroup(matrix, row-1, col, color, checkedCoords, coords) // down
	fillCoordsOfColorGroup(matrix, row+1, col, color, checkedCoords, coords) // up
	fillCoordsOfColorGroup(matrix, row, col-1, color, checkedCoords, coords) // left
	fillCoordsOfColorGroup(matrix, row, col+1, color, checkedCoords, coords) // right
}
