package matrix

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

type testcase struct {
	matrixA [][]int
	matrixB [][]int
	matrixC [][]int
}

var testcases = []testcase{
	{
		matrixA: [][]int{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		},
		matrixB: [][]int{
			{1, 2, 3, 4},
			{2, 2, 3, 4},
			{3, 3, 3, 4},
			{4, 4, 4, 4},
		},
		matrixC: [][]int{
			{30, 31, 34, 40},
			{30, 31, 34, 40},
			{30, 31, 34, 40},
			{30, 31, 34, 40},
		},
	},
}

func TestNaiveMultiplication(t *testing.T) {
	for _, testcase := range testcases {
		C := MultiplyMatrixWithIteration(testcase.matrixA, testcase.matrixB)
		fmt.Println(C)
		assert.Equal(t, testcase.matrixC, C, "Output not equal.")
	}
}
