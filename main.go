package main

import (
	"context"
	"os"
)

func main() {
	go slowServer()
	go errServer()

	ctx := context.Background()
	callBoth(ctx, os.Args[1])
}
