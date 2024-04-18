package main

import "log"

type ringBuffer struct {
	in  chan int
	out chan int
}

func newRingBuffer(inCh chan int, outCh chan int) *ringBuffer {
	return &ringBuffer{
		in:  inCh,
		out: outCh,
	}
}

func (r *ringBuffer) Run() {
	defer close(r.out)
	for v := range r.in {
		select {
		case r.out <- v:
		default:
			<-r.out
			r.out <- v
		}
	}
}

func main() {
	in := make(chan int)
	out := make(chan int, 4) //nolint:gomnd // it's learning code

	rb := newRingBuffer(in, out)

	go rb.Run()

	for i := range 10 {
		in <- i
	}

	close(in)

	for res := range out {
		log.Println(res)
	}
}
