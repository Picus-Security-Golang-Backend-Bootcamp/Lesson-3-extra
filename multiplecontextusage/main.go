package main

import (
	"context"
	"fmt"
	"time"
)

// LongProcess refers to a long network request
func LongProcess(ctx context.Context, duration time.Duration, msg string) error {
	c1 := make(chan string, 1)
	go func() {
		// Simulate processing
		time.Sleep(duration)
		c1 <- msg
	}()

	select {
	case m := <-c1:
		fmt.Println(m)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx := context.Background()
	t := 2 * time.Second

	ctx, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	// Simulate a 2 second process time
	err := LongProcess(ctx, 2*time.Second, "first process")
	fmt.Println(err)

	// Reusing the context.
	s, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	// Simulate a 1 second process time
	err = LongProcess(s, 1*time.Second, "second process")
	fmt.Println(err) // context deadline exceeded
}
