package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type NetworkSender struct{}

func (s *NetworkSender) Write(p []byte) (int, error) {
	return fmt.Printf("Network send: %s", string(p)) //nolint:wrapcheck,forbidigo // it's learning code
}

type SampleInterface interface {
	Method() string
}

type OtherInterface interface {
	SampleInterface
	OtherMethod()
}

type BidirectionalCommunication struct{}

func (b *BidirectionalCommunication) Read(_ []byte) (int, error) {
	return 0, nil
}

func (b *BidirectionalCommunication) Write(_ []byte) (int, error) {
	return 0, nil
}

type SampleStruct struct{}

func (s *SampleStruct) Method() string {
	return ""
}

func main() {
	var w io.Writer

	w = os.Stdout
	_, _ = w.Write([]byte("Hello, World!\n")) // Hello, World!

	w = &NetworkSender{}
	_, _ = w.Write([]byte("Hello, World!\n")) // Network send: Hello, World!

	// type assertion
	var i any = "hello"
	s := i.(string) //nolint:errcheck // it's learning code
	log.Println(s)

	// number := i.(int) // panic: interface conversion, interface{} is string, not int
	number, ok := i.(int)
	if !ok {
		log.Printf("Type %T is not a number, number =%d\n", i, number)
	} else {
		log.Println(number)
	}

	// type conversion
	intVar := string(97)
	log.Println(intVar) // a

	var int32Var int32 = 100500
	var int64Var int64 //nolint:gosimple // it's learning code

	int64Var = int64(int32Var)
	log.Println(int64Var) // 100500

	// type switch
	var t any = "hello"
	switch tNonInterface := t.(type) {
	case string:
		log.Printf("string: %s\n", tNonInterface) // string: hello
	case bool:
		log.Printf("boolean: %t\n", tNonInterface)
	case int:
		log.Printf("integer: %d\n", tNonInterface)
	default:
		log.Printf("unexpected: %T\n", tNonInterface)
	}

	// zero value
	var sample SampleInterface
	log.Println(sample, sample == nil) // <nil> true

	var sampleInstance *SampleStruct = nil //nolint:revive // it's learning code
	sample = sampleInstance
	log.Println(sample, sample == nil) //nolint:staticcheck // <nil> false

	var x io.ReadWriter = &BidirectionalCommunication{}
	_, _ = x.Read([]byte{})
	_, _ = x.Write([]byte{})

	// var y OtherInterface
	// y.Method()      //nolint:govet // panic: runtime error: invalid address or nil pointer deference
	// y.OtherMethod() //nolint:govet // it's learning code
}
