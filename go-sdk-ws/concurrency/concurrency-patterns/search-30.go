package main

import (
	"concurrency-patterns/google"
	"fmt"
	"time"
)

func main() {
	web1, web2, web3 := google.FakeSearch("web1"), google.FakeSearch("web2"), google.FakeSearch("web3")
	image1, image2, image3 := google.FakeSearch("image1"), google.FakeSearch("image2"), google.FakeSearch("image3")
	video1, video2, video3 := google.FakeSearch("video1"), google.FakeSearch("video2"), google.FakeSearch("video3")
	c := make(chan google.Result)

	go func() { c <- google.First("Web", web1, web2, web3) }()
	go func() { c <- google.First("Image", image1, image2, image3) }()
	go func() { c <- google.First("Video", video1, video2, video3) }()
	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}

	return
}
