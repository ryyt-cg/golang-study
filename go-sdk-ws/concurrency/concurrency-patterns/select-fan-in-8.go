package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn8(boring8("joe"), boring8("ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Printf("You're both boring; I'm leaving.")
}

func fanIn8(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func boring8(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller
}
