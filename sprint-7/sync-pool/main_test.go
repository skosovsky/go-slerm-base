package main

import "testing"

// go test -bench=. -benchmem
// BenchmarkWithoutPool-10            13646             84177 ns/op           80000 B/op      10000 allocs/op
// BenchmarkWithPool-10               14659             82250 ns/op               0 B/op          0 allocs/op

func BenchmarkWithoutPool(b *testing.B) {
	var person *Person
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		for range 10000 {
			person = new(Person)
			person.Age = 23 //nolint:govet // it's test
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var person *Person
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		for range 10000 {
			person = personalPool.Get().(*Person) //nolint:errcheck // it's test
			person.Age = 23
			personalPool.Put(person)
		}
	}
}
