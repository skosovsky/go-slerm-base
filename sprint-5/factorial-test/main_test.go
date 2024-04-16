package main

import "testing"

func TestFactorial(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name   string
		num    int
		result int
	}{
		{name: "1", num: 0, result: 1},
		{name: "1", num: 1, result: 1},
		{name: "1", num: 2, result: 2},
		{name: "1", num: 3, result: 6},
		{name: "1", num: 4, result: 24},
		{name: "1", num: 5, result: 120},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			result := factorial(testCase.num)
			if result != testCase.result {
				t.Errorf("Expected %d, got %d", testCase.result, result)
			}
		})
	}
}
