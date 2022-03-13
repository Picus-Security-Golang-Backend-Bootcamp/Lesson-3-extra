package main

import "fmt"

func main() {
	doWork(nil)
	// Bir ton yapılacak iş daha
	fmt.Println("Done.")
}

func doWork(strings <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doWork exited.")
		defer close(completed)
		for s := range strings {
			// Do something interesting
			fmt.Println(s)
		}
	}()
	return completed
}
