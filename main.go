package main

import (
	"context"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go slowServer(&wg)
	go errServer(&wg)
	wg.Wait()

	ctx := context.Background()
	callBoth(ctx, os.Args[1])
}
