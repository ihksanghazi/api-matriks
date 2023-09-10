package services

import (
	"fmt"
)

func AddMatrix(a [][]int, b [][]int) [][]int {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		panic("Matrix dimensions do not match")
	}

	rows := len(a)
	cols := len(a[0])

	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}

	return result
}

func SubtractMatrix(a [][]int, b [][]int) [][]int {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		panic("Matrix dimensions do not match")
	}

	rows := len(a)
	cols := len(a[0])

	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}

	return result
}

func MultiplyMatrix(a [][]int, b [][]int) [][]int {
	rowsA := len(a)
	colsA := len(a[0])
	rowsB := len(b)
	colsB := len(b[0])

	if colsA != rowsB {
		panic("Matrix dimensions do not match")
	}

	result := make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result
}

func TransposeMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	// Membuat matriks baru dengan baris dan kolom yang dibalik
	transposed := make([][]float64, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]float64, rows)
	}

	// Melakukan transposisi
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func InverseMatriks(matrix [][]float64) ([][]float64, error) {
	if len(matrix) != 2 || len(matrix[0]) != 2 || len(matrix[1]) != 2 {
		return nil, fmt.Errorf("matriks must be 2x2")
	}

	a := matrix[0][0]
	b := matrix[0][1]
	c := matrix[1][0]
	d := matrix[1][1]

	det := a*d - b*c

	if det == 0 {
		return nil, fmt.Errorf("matriks tidak memiliki invers karena determinannya nol")
	}

	inverseMatrix := make([][]float64, 2)
	inverseMatrix[0] = []float64{d / det, -b / det}
	inverseMatrix[1] = []float64{-c / det, a / det}

	return inverseMatrix, nil
}

func DeterminantMatriks(arr [][]float64) float64 {
	determinant := (arr[0][0] * arr[1][1]) - (arr[0][1] * arr[1][0])
	return determinant
}

func ReduceMatrix(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	// Inisialisasi indeks baris dan kolom
	i := 0
	j := 0

	for i < rows && j < cols {
		// Temukan baris dengan elemen terbesar pada kolom ke-j
		maxRow := i
		for k := i + 1; k < rows; k++ {
			if matrix[k][j] > matrix[maxRow][j] {
				maxRow = k
			}
		}

		// Tukar baris terbesar dengan baris saat ini (baris ke-i)
		for k := j; k < cols; k++ {
			matrix[i][k], matrix[maxRow][k] = matrix[maxRow][k], matrix[i][k]
		}

		// Membuat elemen diagonal menjadi 1
		pivot := matrix[i][j]
		for k := j; k < cols; k++ {
			matrix[i][k] /= pivot
		}

		// Mengurangkan baris lain dari baris saat ini
		for k := 0; k < rows; k++ {
			if k != i {
				factor := matrix[k][j]
				for l := j; l < cols; l++ {
					matrix[k][l] -= factor * matrix[i][l]
				}
			}
		}

		i++
		j++
	}

	return matrix
}

func CreateIdentityMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func CreateDiagonalMatrix(size, diagonalValue int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = diagonalValue
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
