// Custom struct which implements multiple MaxSubArray methods
package max_sub_array

import (
	"errors"
	"math"
)

type CustomArray struct {
	Array      []int
	Operations int
}

var (
	ErrTooManyMode = errors.New("too many parameters provided")
)

func (arr *CustomArray) GetMaxSubArray(mode ...string) (start, end, sum int, err error) {
	if len(mode) > 1 {
		return 0, 0, 0, ErrTooManyMode
	}

	if mode[0] == "naive" {
		start, end, sum = arr.maxSubArrayNaive()
	} else if mode[0] == "recurse" {
		start, end, sum = arr.maxSubArrayDivideAndConquer()
	} else {
		start, end, sum = arr.maxSubArrayLinear()
	}

	return start, end, sum, nil
}

// Iterate through every single subarray in A, and keep track of the subarray
// with the maximum sum. Return the start and end index of the maximum subarray,
// along with the sum.
//
// Returns (start int, end int, sum int)
func (arr *CustomArray) maxSubArrayNaive() (int, int, int) {
	var sum, start, end int

	sum = int(math.Inf(-1))
	var currentSum int
	for i := range arr.Array {
		arr.incrementOperationCount()
		currentSum = arr.Array[i]

		// To check if the last element is the maximum subarray when i is the
		// last element
		if currentSum >= sum {
			sum = currentSum
			start = i
			end = i
		}

		for j := i + 1; j < len(arr.Array); j++ {
			arr.incrementOperationCount()

			currentSum += arr.Array[j]
			if currentSum >= sum {
				sum = currentSum
				start = i
				end = j
			}
		}
	}
	return start, end, sum
}

// Driver function which calls maxSubArrayRecurse. maxSubArrayRecurse
// implements finding max sub array with divide and conquer technique.
func (arr *CustomArray) maxSubArrayDivideAndConquer() (int, int, int) {
	low := 0
	high := len(arr.Array) - 1

	start, end, sum := maxSubArrayRecurse(arr, low, high)
	return start, end, sum
}

// Divide and Conquery technique for finding maximum subarray
//
// returns (star int, end int, sum int)
func maxSubArrayRecurse(a *CustomArray, start, end int) (int, int, int) {
	// base case
	if start == end {
		a.incrementOperationCount()
		return start, end, a.Array[start]
	}

	// divide - solving recurrence cases
	mid := (start + end) / 2
	a.incrementOperationCount()

	left_start, left_end, left_sum := maxSubArrayRecurse(a, start, mid)

	right_start, right_end, right_sum := maxSubArrayRecurse(a, mid+1, end)

	cross_start, cross_end, cross_sum := maxCrossingSubArray(a, start, mid, end)

	if left_sum >= right_sum && left_sum >= cross_sum {
		return left_start, left_end, left_sum
	} else if right_sum >= left_sum && right_sum >= cross_sum {
		return right_start, right_end, right_sum
	} else {
		return cross_start, cross_end, cross_sum
	}
}

func maxCrossingSubArray(a *CustomArray, start, mid, end int) (int, int, int) {
	// leftMax and rightMax are the index containing the crossing subarray
	var leftMax, rightMax, totalSum int

	leftSum := int(math.Inf(-1))
	leftCurrentSum := 0
	// edge case, mid and start overlap
	if start == mid {
		a.incrementOperationCount()
		leftSum = a.Array[start]
		leftMax = start
	} else {
		// go through left side from the middle
		for i := mid; i >= start; i-- {
			a.incrementOperationCount()
			leftCurrentSum += a.Array[i]
			if leftCurrentSum > leftSum {
				leftSum = leftCurrentSum
				leftMax = i
			}
		}
	}

	rightSum := int(math.Inf(-1))
	rightCurrentSum := 0
	// edge case, mid and end overlap
	if end == mid {
		a.incrementOperationCount()
		rightSum = a.Array[end]
		rightMax = end
	} else {
		// go through right side from the middle
		for j := mid + 1; j <= end; j++ {
			a.incrementOperationCount()
			rightCurrentSum += a.Array[j]
			if rightCurrentSum > rightSum {
				rightSum = rightCurrentSum
				rightMax = j
			}
		}
	}

	totalSum = leftSum + rightSum
	return leftMax, rightMax, totalSum
}

// Linear solution.
//
// returns (start, end, sum)
func (arr *CustomArray) maxSubArrayLinear() (int, int, int) {
	sum := int(math.Inf(-1))
	current := 0
	currentStart := 0
	start := 0
	end := 0

	for i := range arr.Array {
		arr.incrementOperationCount()
		current = current + arr.Array[i]

		// If the current element is larger than the current sum, it means that
		// the previous sum must be negative, so we can safely discard the previous
		// subarray, and the current element could be the start of the actual max
		// subarray.
		if arr.Array[i] > current {
			currentStart = i
			current = arr.Array[i]
		}

		// When the current sum is larger than existing sum, we will replace the max
		// sum value, and the existing indices
		if current > sum {
			sum = current
			start = currentStart
			end = i
		}
	}
	return start, end, sum
}

// Increment the number of operation by 1
func (arr *CustomArray) incrementOperationCount() {
	arr.Operations += 1
}
