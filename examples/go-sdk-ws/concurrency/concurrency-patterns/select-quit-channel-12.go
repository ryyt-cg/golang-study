package main

import (
	"fmt"
	"math/rand"
)

// timeout for the whole conversation
func main() {
	quit := make(chan string)
	c := boring12("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring12(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				//cleanup()
				quit <- "See you!"
				return
			}
		}
	}()

	return c // return the channel to the caller
}
