package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// A channel provides a connection between two goroutines, allowing them to communicate.
	c := make(chan string)
	go boring2("boring!", c)

	// main function is listening to the channel
	// <-c is a blocking operation
	// main function will wait until a message is received
	// it loops 5 times to receive 5 messages and exit
	for i := 0; i < 5; i++ {
		// receive a message from the channel and print it.
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

// boring2
// The function sends a generated message to the channel and sleeps for a random duration.
func boring2(msg string, c chan<- string) {
	for i := 0; ; i++ {
		// send a message to the channel, blocking until the message is received
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
