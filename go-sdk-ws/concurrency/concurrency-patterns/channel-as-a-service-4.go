package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring4("joe")
	ann := boring4("ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-ann) // blocking operation, waiting for the value to be sent
		fmt.Println(<-joe)
	}

	fmt.Printf("You're both boring; I'm leaving.")
}

func boring4(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // blocking operation, waiting for the value to be received
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller
}
