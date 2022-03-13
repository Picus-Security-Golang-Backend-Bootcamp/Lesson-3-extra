package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan interface{})
	randStream := generateNumber(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}

func generateNumber(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	rand.Seed(time.Now().Unix())
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				return
			}
		}
	}()

	return randStream
}
