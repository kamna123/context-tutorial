package main

import (
	"context"
	"fmt"
	"time"
)

func ScheduleTask(ctx context.Context, taskName string) {

	select {
	case <-time.After(4 * time.Second):
		fmt.Printf("Task '%s' completed\n", taskName)
	case <-ctx.Done():
		fmt.Printf("Task '%s' canceled: %v\n", taskName, ctx.Err())
	}
}

func main() {
	parentCtx := context.Background()
	deadline := time.Now().Add(1 * time.Second)
	childCtx1, cancel := context.WithDeadline(parentCtx, deadline)
	defer cancel()

	go ScheduleTask(childCtx1, "Data Processing")

	childChildCtx1, cancel := context.WithDeadline(childCtx1, time.Now().Add(100*time.Second))
	defer cancel()
	go ScheduleTask(childChildCtx1, "File Processing")

	childCtx2, cancel := context.WithDeadline(parentCtx, time.Now().Add(10*time.Second))
	defer cancel()
	go ScheduleTask(childCtx2, "Some remote Processing")
	// ...
	// Perform other concurrent operations
	// ...
	time.Sleep(10 * time.Second)
}
