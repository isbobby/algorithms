package max_sub_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	inputArray    []int
	expectedStart int
	expectedEnd   int
	expectedSum   int
	expectedErr   error
}

var testcases = []testcase{
	{
		inputArray:    []int{1, -3, 4, 1, -1},
		expectedStart: 2,
		expectedEnd:   3,
		expectedSum:   5,
		expectedErr:   nil,
	},
	{
		inputArray:    []int{1, -2, 1, 1, 1},
		expectedStart: 2,
		expectedEnd:   4,
		expectedSum:   3,
		expectedErr:   nil,
	},
	{
		inputArray:    []int{-10, -10, -10, -10, -1},
		expectedStart: 4,
		expectedEnd:   4,
		expectedSum:   -1,
		expectedErr:   nil,
	},
	{
		inputArray:    []int{-10, -10, -10, -10, 2, 3, -2, 5, -9, 3, 1, -3, 4, 1, -1},
		expectedStart: 4,
		expectedEnd:   7,
		expectedSum:   8,
		expectedErr:   nil,
	},
}

func runTests(mode string, t *testing.T) {
	for i, tc := range testcases {
		// initialise custom array
		fmt.Printf("Running test case %d\n", i)
		arr := CustomArray{Array: tc.inputArray, Operations: 0}

		outputStart, outputEnd, outputSum, err := arr.GetMaxSubArray(mode)

		assert.Equal(t, tc.expectedStart, outputStart, "Start index not equal.")
		assert.Equal(t, tc.expectedEnd, outputEnd, "End index not equal.")
		assert.Equal(t, tc.expectedSum, outputSum, "Subarray sum not equal.")
		assert.Equal(t, tc.expectedErr, err, "Error not equal.")

		fmt.Printf("Total number of operations for case %d with %d elements is %d\n", i, len(tc.inputArray), arr.Operations)
	}
}

func TestMaxArrayNaive(t *testing.T) {
	runTests("naive", t)
}

func TestMaxArrayDivideAndConquer(t *testing.T) {
	runTests("recurse", t)
}

func TestMaxArrayLinear(t *testing.T) {
	runTests("linear", t)
}
