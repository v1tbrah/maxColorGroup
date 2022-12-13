package matrix

import (
	"strconv"
	"testing"
)

func Test_GetCoordsOfMaxColorGroup(t *testing.T) {
	tests := []struct {
		name                      string
		description               string
		representation            string
		matrix                    [][]int
		wantCoordsOfMaxColorGroup []Coord
		wantColor                 int
	}{
		{
			name: "empty matrix",
			matrix: func() [][]int {
				return nil
			}(),
		},
		{
			name: "one color matrix",
			matrix: func() [][]int {
				matrix := make([][]int, 3)
				for row, _ := range matrix {
					matrix[row] = make([]int, 3)
				}
				for row, _ := range matrix {
					for col, _ := range matrix[row] {
						matrix[row][col] = 0
					}
				}
				return matrix
			}(),
			wantCoordsOfMaxColorGroup: []Coord{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},
		{
			name: "three color matrix",
			matrix: func() [][]int {
				// 0 0 0
				// 1 0 1
				// 2 1 2
				matrix := make([][]int, 3)
				for row, _ := range matrix {
					matrix[row] = make([]int, 3)
				}
				matrix[0][0] = 0
				matrix[0][1] = 0
				matrix[0][2] = 0
				matrix[1][0] = 1
				matrix[1][1] = 0
				matrix[1][2] = 1
				matrix[2][0] = 2
				matrix[2][1] = 1
				matrix[2][2] = 2

				return matrix
			}(),
			wantCoordsOfMaxColorGroup: []Coord{
				{0, 0}, {0, 1}, {0, 2},
				{1, 1},
			},
		},
		{
			name: "three color matrix with check diagonal",
			matrix: func() [][]int {
				// 0 0 0
				// 1 0 1
				// 0 1 2
				matrix := make([][]int, 3)
				for row, _ := range matrix {
					matrix[row] = make([]int, 3)
				}
				matrix[0][0] = 0
				matrix[0][1] = 0
				matrix[0][2] = 0
				matrix[1][0] = 1
				matrix[1][1] = 0
				matrix[1][2] = 1
				matrix[2][0] = 0
				matrix[2][1] = 1
				matrix[2][2] = 2

				return matrix
			}(),
			wantCoordsOfMaxColorGroup: []Coord{
				{0, 0}, {0, 1}, {0, 2},
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCoordsOfMaxColorGroup, gotColor := GetCoordsOfMaxColorGroup(tt.matrix)

			if gotColor != tt.wantColor {
				t.Errorf("getCoordsOfMaxColorGroup() gotColor = %v, want %v", gotColor, tt.wantColor)
			}
			if !equalWithoutOrder(gotCoordsOfMaxColorGroup, tt.wantCoordsOfMaxColorGroup) {
				t.Errorf("getCoordsOfMaxColorGroup() gotCoordsOfMaxColorGroup = %v, want %v", gotCoordsOfMaxColorGroup, tt.wantCoordsOfMaxColorGroup)
			}
		})
	}
}

func equalWithoutOrder(first, second []Coord) bool {
	if len(first) != len(second) {
		return false
	}

	hashFirst := make(map[string]struct{}, len(first))
	for _, coord := range first {
		key := strconv.Itoa(coord.X) + (strconv.Itoa(coord.Y))
		hashFirst[key] = struct{}{}
	}

	for _, coord := range second {
		key := strconv.Itoa(coord.X) + (strconv.Itoa(coord.Y))
		if _, ok := hashFirst[key]; !ok {
			return false
		}
	}

	return true
}
