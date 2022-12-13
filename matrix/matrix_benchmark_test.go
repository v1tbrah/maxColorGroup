package matrix

import (
	"math/rand"
	"testing"
)

func BenchmarkGetCoordsOfMaxColorGroup_Matrix100(b *testing.B) {

	mat := make([][]int, 10)
	for row, _ := range mat {
		mat[row] = make([]int, 10)
	}

	for row, _ := range mat {
		for col, _ := range mat[row] {
			mat[row][col] = rand.Intn(10)
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetCoordsOfMaxColorGroup(mat)
	}
}

func BenchmarkGetCoordsOfMaxColorGroup_Matrix10000(b *testing.B) {

	mat := make([][]int, 100)
	for row, _ := range mat {
		mat[row] = make([]int, 100)
	}

	for row, _ := range mat {
		for col, _ := range mat[row] {
			mat[row][col] = rand.Intn(10)
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetCoordsOfMaxColorGroup(mat)
	}
}

func BenchmarkGetCoordsOfMaxColorGroup_Matrix1000000(b *testing.B) {

	mat := make([][]int, 1000)
	for row, _ := range mat {
		mat[row] = make([]int, 1000)
	}

	for row, _ := range mat {
		for col, _ := range mat[row] {
			mat[row][col] = rand.Intn(10)
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GetCoordsOfMaxColorGroup(mat)
	}
}
