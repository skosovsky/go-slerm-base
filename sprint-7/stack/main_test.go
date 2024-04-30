package main

import (
	"testing"

	"github.com/google/uuid"
)

// go test -bench=. -benchmem
// BenchmarkBigStackByValue_Push-10                12275085                97.44 ns/op          373 B/op          0 allocs/op
// BenchmarkBigStackByPointer_Push-10              33829690                36.84 ns/op           47 B/op          0 allocs/op
// BenchmarkSmallStackByValue_Push-10              218043565                5.378 ns/op          86 B/op          0 allocs/op
// BenchmarkSmallStackByPointer_Push-10            43724535                25.33 ns/op           46 B/op          0 allocs/op

func getBigStackEntry() BigStackEntry {
	return BigStackEntry{
		valueInt:     123,
		valueFloat:   456.789,
		payload:      uuid.NewString(),
		otherPayload: uuid.NewString(),
		embeddedParams: struct {
			A int
			B int
			C int
		}{A: 10, B: 20, C: 30},
	}
}

func getSmallStackEntry() SmallStackEntry {
	return SmallStackEntry{
		valueInt:   123,
		valueFloat: 456.789,
	}
}

func BenchmarkBigStackByValue_Push(b *testing.B) {
	el := getBigStackEntry()
	stack := BigStackByValue{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(el)
	}
}

func BenchmarkBigStackByPointer_Push(b *testing.B) {
	elValue := getBigStackEntry()
	el := &elValue
	stack := BigStackByPointer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(el)
	}
}

func BenchmarkSmallStackByValue_Push(b *testing.B) {
	el := getSmallStackEntry()
	stack := SmallStackByValue{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(el)
	}
}

func BenchmarkSmallStackByPointer_Push(b *testing.B) {
	elValue := getSmallStackEntry()
	el := &elValue
	stack := SmallStackByPointer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(el)
	}
}
