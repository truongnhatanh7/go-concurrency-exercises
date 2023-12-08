//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

var msg = make(chan *Tweet)
var done = make(chan bool)

func producer(stream Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			done <- true
			return
		}
		msg <- tweet
	}
}

func consumer() {
	for t := range msg {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	go producer(stream)
	go consumer()
	<-done

	fmt.Printf("Process took %s\n", time.Since(start))
}
