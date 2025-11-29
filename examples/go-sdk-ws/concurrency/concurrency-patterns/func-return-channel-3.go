package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	ch := boring3("boring!") // function returning a channel
	for i := 0; i < 5; i++ {
		// <-ch is a blocking operation
		fmt.Printf("You say: %q\n", <-ch)
	}

	fmt.Printf("You're boring; I'm leaving.\n")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s\n", elapsed)
}

// boring3
// The function returns a receiver channel to the caller.
func boring3(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller
}
