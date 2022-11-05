package matrix

func MultiplyMatrixWithIteration(A [][]int, B [][]int) [][]int {
	n := len(A)    // number of rows of matrix A
	m := len(B[0]) // number of cols of matrix B

	res := [][]int{}

	for i := 0; i < n; i++ {
		row := []int{}

		for j := 0; j < m; j++ {

			// sum of C[i][j]
			sum := 0
			for k := 0; k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			// append to row
			row = append(row, sum)
		}
		res = append(res, row)
	}
	return res
}
