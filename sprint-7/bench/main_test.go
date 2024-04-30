package main

import "testing"

func BenchmarkName(b *testing.B) {
	// expensiveSetup() prepare test
	b.ResetTimer() // start timer now

	for range b.N {
		b.StopTimer() // pause timer now
		// addExpensiveSetup() additional prepare
		b.StartTimer() // resume timer now
		// functionUnderTest() no testing func
	}
}
