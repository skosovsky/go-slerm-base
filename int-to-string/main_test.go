package main

import "testing"

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 {
			itoa(v)
		}
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 {
			sprint(v)
		}
	}
}

func BenchmarkItoaPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 {
			itoaPlus(v)
		}
	}
}

func BenchmarkAppendInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for v := range 1000000 {
			appendInt(v)
		}
	}
}
