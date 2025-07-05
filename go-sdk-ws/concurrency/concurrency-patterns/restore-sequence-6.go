package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Message
// Restoring sequencing using fan-in pattern
// Send a channel on a channel, making goroutines wait for its turn
// Receive all messages, then enable them again by sending on a private channel
// First we define a message type that contains a channel for the reply
type Message struct {
	str  string
	wait chan bool // signal channel
}

func main() {
	c := fanIn2(boring6("joe"), boring6("ann"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true
	}

	fmt.Printf("You're both boring; I'm leaving.")
}

func fanIn2(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)

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

func boring6(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) // shared between all messages

	go func() {
		for i := 0; ; i++ {
			message := Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			c <- message
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt // wait for the main function to signal
		}
	}()

	return c // return the channel to the caller
}
