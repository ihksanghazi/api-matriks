package services

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
