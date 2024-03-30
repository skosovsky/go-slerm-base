package main

import "testing"

func BenchmarkSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		small()
	}
}

func BenchmarkBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		big()
	}
}
