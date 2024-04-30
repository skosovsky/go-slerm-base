package main

import "sync"

type Person struct {
	Age int
}

var personalPool = sync.Pool{ //nolint:gochecknoglobals // it's example
	New: func() any { return new(Person) },
}

func main() {

}
