package main

import "log"

// bad design
// package event_emitter

type BadEventEmitterIface interface {
	Emit(event string)
}

type BadEventEmitter struct{}

func NewBadEventEmitter() BadEventEmitterIface { //nolint:ireturn // it's bad design
	return &BadEventEmitter{}
}

func (b BadEventEmitter) Emit(event string) {
	log.Println("emit event", event)
}

// package handler

func BadHandler(e BadEventEmitterIface) {
	e.Emit("login")
}

// good design
// package event_emitter

type EventEmitter struct{}

func NewGoodEventEmitter() EventEmitter {
	return EventEmitter{}
}

func (e EventEmitter) Emit(event string) {
	log.Println("emit event", event)
}

// package handler

type GoodEventEmitterIface interface {
	Emit(event string)
}

func GoodHandler(e GoodEventEmitterIface) {
	e.Emit("login")
}

func acceptInterfaceReturnStructs() {
	GoodHandler(NewGoodEventEmitter()) // emit event login
}

func main() {
	acceptInterfaceReturnStructs()
}
