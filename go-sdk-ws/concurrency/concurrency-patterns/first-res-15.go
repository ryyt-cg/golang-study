package main

import (
	"fmt"
	"math/rand"
	"rhtran/go-study/07-concurrency/3-concurrency-patterns/google"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	result := google.First("golang",
		google.FakeSearch("replica 1"),
		google.FakeSearch("replica 2"),
	)
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}
