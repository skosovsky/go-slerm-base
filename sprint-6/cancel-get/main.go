package main

import (
	"golang.org/x/net/context"
	"log"
)

func main() {
	ctx := context.Background()
	log.Println(ctx)
}
