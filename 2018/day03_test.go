package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day03()

	assert.Equal(118840, part1)
	assert.Equal(919, part2)
}

var benchResult int

func BenchmarkXY(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create matrix
		fabicMatrix := make([][]int, 1000)
		for i := range fabicMatrix {
			fabicMatrix[i] = make([]int, 1000)
		}

		for _, claim := range claimList {
			for x := claim.left; x < claim.left+claim.width; x++ {
				for y := claim.top; y < claim.top+claim.height; y++ {
					fabicMatrix[x][y]++
				}
			}
		}

		overlappingSquares := 0
		for x, a := range fabicMatrix {
			for y := range a {
				if fabicMatrix[x][y] > 1 {
					overlappingSquares++
				}
			}
		}

		benchResult = overlappingSquares
	}
}

func BenchmarkXYint8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create matrix
		fabicMatrix := make([][]int8, 1000)
		for i := range fabicMatrix {
			fabicMatrix[i] = make([]int8, 1000)
		}

		for _, claim := range claimList {
			for x := claim.left; x < claim.left+claim.width; x++ {
				for y := claim.top; y < claim.top+claim.height; y++ {
					fabicMatrix[x][y]++
				}
			}
		}

		overlappingSquares := 0
		for x, a := range fabicMatrix {
			for y := range a {
				if fabicMatrix[x][y] > 1 {
					overlappingSquares++
				}
			}
		}
		benchResult = overlappingSquares
	}
}
func BenchmarkXYuint8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create matrix
		fabicMatrix := make([][]uint8, 1000)
		for i := range fabicMatrix {
			fabicMatrix[i] = make([]uint8, 1000)
		}

		for _, claim := range claimList {
			for x := claim.left; x < claim.left+claim.width; x++ {
				for y := claim.top; y < claim.top+claim.height; y++ {
					fabicMatrix[x][y]++
				}
			}
		}

		overlappingSquares := 0
		for x, a := range fabicMatrix {
			for y := range a {
				if fabicMatrix[x][y] > 1 {
					overlappingSquares++
				}
			}
		}
		benchResult = overlappingSquares
	}
}
func BenchmarkXYNoAlloc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create matrix
		fabicMatrix := [1000][1000]int{}

		for _, claim := range claimList {
			for x := claim.left; x < claim.left+claim.width; x++ {
				for y := claim.top; y < claim.top+claim.height; y++ {
					fabicMatrix[x][y]++
				}
			}
		}

		overlappingSquares := 0
		for x, a := range fabicMatrix {
			for y := range a {
				if fabicMatrix[x][y] > 1 {
					overlappingSquares++
				}
			}
		}
		benchResult = overlappingSquares
	}
}

func BenchmarkXYNoAllocInt8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create matrix
		fabicMatrix := [1000][1000]uint8{}

		for _, claim := range claimList {
			for x := claim.left; x < claim.left+claim.width; x++ {
				for y := claim.top; y < claim.top+claim.height; y++ {
					fabicMatrix[x][y]++
				}
			}
		}

		overlappingSquares := 0
		for x, a := range fabicMatrix {
			for y := range a {
				if fabicMatrix[x][y] > 1 {
					overlappingSquares++
				}
			}
		}
		benchResult = overlappingSquares
	}
}

func BenchmarkYX(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Create matrix
		fabicMatrix := make([][]int, 1000)
		for i := range fabicMatrix {
			fabicMatrix[i] = make([]int, 1000)
		}

		for _, claim := range claimList {
			for x := claim.left; x < claim.left+claim.width; x++ {
				for y := claim.top; y < claim.top+claim.height; y++ {
					fabicMatrix[y][x]++
				}
			}
		}

		overlappingSquares := 0
		for x, a := range fabicMatrix {
			for y := range a {
				if fabicMatrix[y][x] > 1 {
					overlappingSquares++
				}
			}
		}
		benchResult = overlappingSquares
	}
}
