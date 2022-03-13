package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randStream := generateNumber()
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func generateNumber() <-chan int {
	randStream := make(chan int)
	rand.Seed(time.Now().Unix())
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			randStream <- rand.Intn(1000)
		}
	}()

	return randStream
}
