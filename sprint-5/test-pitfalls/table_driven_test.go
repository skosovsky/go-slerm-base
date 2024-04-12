package main

import "testing"

// go test -v ./table_driven_test.go
// === RUN   TestLog
// === PAUSE TestLog
// === CONT  TestLog
// === RUN   TestLog/test_1
// table_driven_test.go:18: 1
// === RUN   TestLog/test_2
// table_driven_test.go:18: 2
// === RUN   TestLog/test_3
// table_driven_test.go:18: 3
// === RUN   TestLog/test_4
// table_driven_test.go:18: 4
// --- PASS: TestLog (0.00s)
// --- PASS: TestLog/test_1 (0.00s)
// --- PASS: TestLog/test_2 (0.00s)
// --- PASS: TestLog/test_3 (0.00s)
// --- PASS: TestLog/test_4 (0.00s)
// PASS
// ok      command-line-arguments  0.362s

// go test -v -run TestTLog/test_1 ./table_driven_test.go
// === RUN   TestTLog
// === PAUSE TestTLog
// === CONT  TestTLog
// === RUN   TestTLog/test_1
// table_driven_test.go:39: 1
// --- PASS: TestTLog (0.00s)
// --- PASS: TestTLog/test_1 (0.00s)
// PASS
// ok      command-line-arguments  0.285s

func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.value)
		})
	}
}
