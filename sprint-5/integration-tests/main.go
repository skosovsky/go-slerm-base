package main

import (
	"context"

	"github.com/skosovsky/go-slerm-base/sprint-5/integration-tests/logic"
)

func main() {
	client := logic.GetClient("localhost:6379")
	ctx := context.Background()
	logic.SetValue(ctx, client, "some value")
}
