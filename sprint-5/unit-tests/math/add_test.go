package math_test

import (
	"testing"

	"github.com/skosovsky/go-slerm-base/sprint-5/unit-tests/math"
)

// go test ./math/
// ok      github.com/skosovsky/go-slerm-base/sprint-5/math-tests/math     0.356s

// go test .
// ?       github.com/skosovsky/go-slerm-base/sprint-5/math-tests  [no test files]

// go test ./...
// ?       github.com/skosovsky/go-slerm-base/sprint-5/math-tests  [no test files]
// ok      github.com/skosovsky/go-slerm-base/sprint-5/math-tests/math     0.173s

// go test ./...
// ?       github.com/skosovsky/go-slerm-base/sprint-5/math-tests  [no test files]
// ok      github.com/skosovsky/go-slerm-base/sprint-5/math-tests/math     (cached)

// go test ./... -count 1 // how count launch test
// ?       github.com/skosovsky/go-slerm-base/sprint-5/math-tests  [no test files]
// ok      github.com/skosovsky/go-slerm-base/sprint-5/math-tests/math     0.282s

// go test -race ./...
// ?       github.com/skosovsky/go-slerm-base/sprint-5/math-tests  [no test files]
// ok      github.com/skosovsky/go-slerm-base/sprint-5/math-tests/math     1.188s

// go test -coverprofile cover.out ./... && go tool cover -html=cover.out -o /tmp/cover.html
// github.com/skosovsky/go-slerm-base/sprint-5/math-tests          coverage: 0.0% of statements
// ok      github.com/skosovsky/go-slerm-base/sprint-5/math-tests/math     0.100s  coverage: 100.0% of statements

func TestAddAlone(t *testing.T) {
	got := math.Add(1, 1)
	expected := 2

	if got != expected {
		t.Fail()
	}
}

func TestAddTableDriver(t *testing.T) {
	testCases := map[string]struct {
		a      int
		b      int
		result int
	}{
		"sum equal": {
			a:      10,
			b:      10,
			result: 20,
		},
		"sum with zero": {
			a:      0,
			b:      15,
			result: 15,
		},
	}

	for _, testCase := range testCases {
		got := math.Add(testCase.a, testCase.b)

		if got != testCase.result {
			t.Errorf("Expected %d, got %d", testCase.result, got)
		}
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{a: 1, b: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
