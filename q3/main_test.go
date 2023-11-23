package main

import (
	"reflect"
	"testing"
)

func TestFindOddInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{7}, 7},
		{[]int{0}, 0},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 1, 0, 1, 0}, 0},
		{[]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 4},
		// Add more test cases here
	}

	for _, test := range tests {
		result := FindOddInt(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
