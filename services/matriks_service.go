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

func TransposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	// Membuat matriks baru dengan baris dan kolom yang dibalik
	transposed := make([][]int, cols)
	for i := 0; i < cols; i++ {
		transposed[i] = make([]int, rows)
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
