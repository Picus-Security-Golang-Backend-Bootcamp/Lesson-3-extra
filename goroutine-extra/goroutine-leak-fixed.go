package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	// Bir ton yapılacak iş daha
	fmt.Println("Done.")
}

func doWork(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(terminated)
		for {
			select {
			case s := <-strings:
				// Do something interesting
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return terminated
}
