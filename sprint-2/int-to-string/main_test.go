package main

import "testing"

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 { //nolint:typecheck
			itoa(v)
		}
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 { //nolint:typecheck
			sprint(v)
		}
	}
}

func BenchmarkItoaPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 { //nolint:typecheck
			itoaPlus(v)
		}
	}
}

func BenchmarkAppendInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 { //nolint:typecheck
			appendInt(v)
		}
	}
}
