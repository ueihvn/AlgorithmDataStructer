package main

import (
	"fmt"
	"math/rand"
	"time"
)

func crawl(v, index int) {
	time.Sleep(time.Second / 100)
	fmt.Printf("WORKER %d:+ FETCH URL %d\n", index+1, v)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	worker := 5
	lenBuffer := 500
	chINURL := make(chan int, lenBuffer)
	quit := make(chan bool)

	lenArrURL := 100
	urls := generateSliceInt(lenArrURL, lenArrURL)

	// send value to channel
	go func() {
		defer close(chINURL)
		for i := 0; i < lenArrURL; i++ {
			chINURL <- urls[i]
		}
	}()

	// start worker
	for i := 0; i < worker; i++ {
		go func(index int) {
			for v := range chINURL {
				crawl(v, index)
			}
			quit <- true
		}(i)
	}

	// quit channel
	for i := 0; i < worker; i++ {
		<-quit
	}

}
