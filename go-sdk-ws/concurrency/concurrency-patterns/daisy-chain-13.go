package main

import (
	"fmt"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	start := time.Now()
	const n = 100_000
	leftmost := make(chan int)

	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)

	fmt.Println("Duration:", time.Since(start))
}
