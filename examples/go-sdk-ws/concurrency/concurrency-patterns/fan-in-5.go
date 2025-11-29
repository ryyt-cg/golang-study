package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring5("joe"), boring5("ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Printf("You're both boring; I'm leaving.")
}

// fanIn function takes two channels as input and returns a single channel
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func boring5(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller
}
