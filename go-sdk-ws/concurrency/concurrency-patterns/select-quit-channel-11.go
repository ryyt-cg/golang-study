package main

import (
	"fmt"
	"math/rand"
	"time"
)

// timeout for the whole conversation
func main() {
	quit := make(chan bool)
	c := boring11("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boring11(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller
}
