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

	checkedCoords := make(map[Coord]struct{}, len(matrix) * len(matrix[0]))
	coordForCheck := Coord{}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			coordForCheck.X = row
			coordForCheck.Y = col
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

func fillCoordsOfColorGroup(matrix [][]int, row, col, color int, checkedCoords map[Coord]struct{}, coords *[]Coord) {
	if len(matrix) == 0 {
		return
	}
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] != color {
		return
	}
	if _, ok := checkedCoords[Coord{row, col}]; ok {
		return
	}
	checkedCoords[Coord{row, col}] = struct{}{}

	*coords = append(*coords, Coord{row, col})

	fillCoordsOfColorGroup(matrix, row - 1, col,  color, checkedCoords, coords) // down
	fillCoordsOfColorGroup(matrix, row + 1, col,  color, checkedCoords, coords) // up
	fillCoordsOfColorGroup(matrix, row, col - 1,  color, checkedCoords, coords) // left
	fillCoordsOfColorGroup(matrix, row, col + 1,  color, checkedCoords, coords) // right
}
