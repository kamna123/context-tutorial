package main

import (
	"context"
	"fmt"
	"time"
)

func makeAPIRequest(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(" API request cancelled ", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("API request completed")
	}
}

func main() {
	parentCtx := context.Background()
	childCtx1, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()

	go makeAPIRequest(childCtx1)
	childCtx2, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()
	go makeAPIRequest(childCtx2)
	// ...
	// Perform other concurrent operations
	// ...
	time.Sleep(5 * time.Second)
}
