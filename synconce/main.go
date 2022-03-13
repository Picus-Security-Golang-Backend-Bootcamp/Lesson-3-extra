package main

import (
	"fmt"
	"sync"
)

func main() {
	//SyncOnce()
	//SyncOnceDeadlock()
	SyncOnceWithMultipleFunction()
}

func SyncOnce() {
	var count int

	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
			increment()
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}

func SyncOnceWithMultipleFunction() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

func SyncOnceDeadlock() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) }

	onceA.Do(initA)
}
