package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var countResult uint32
var wg sync.WaitGroup

// function return slice of int have length = n, and < max
func generateSliceInt(n int, max int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}

func workerJob(ch chan int, quit <-chan int, index int) {
	for {
		select {
		case v, isDead := <-ch:
			if isDead {
				timeFetch := rand.Intn(10)
				time.Sleep(time.Millisecond * time.Duration(timeFetch))
				fmt.Printf("WORKER %d:+ FETCH URL %d TIME: %d ms\n", index+1, v, timeFetch)
				atomic.AddUint32(&countResult, 1)
				// < 5 successfully get data
				// else fail send url back to channel
				// if timeFetch < 8 {
				// 	fmt.Printf("WORKER %d:+ FETCH URL %d TIME: %d ms\n", index+1, v, timeFetch)
				// 	atomic.AddUint32(&countResult, 1)
				// } else {
				// 	ch <- v
				// 	fmt.Printf("WORKER %d:- FETCH URL %d TIME: %d ms\n", index+1, v, timeFetch)
				// }
			}
		case <-quit:
			fmt.Printf("WORKER %d: QUIT\n", index+1)
			wg.Done()
			break
		}
	}
}

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	worker := 5
// 	lenBuffer := 500
// 	chINURL := make(chan int, lenBuffer)

// 	quits := make([]chan int, worker)
// 	for i := range quits {
// 		quits[i] = make(chan int)
// 	}

// 	lenArrURL := 1000
// 	urls := generateSliceInt(lenArrURL, lenArrURL)

// 	// send value to channel
// 	go func() {
// 		defer close(chINURL)
// 		for i := 0; i < lenArrURL; i++ {
// 			chINURL <- urls[i]
// 		}
// 	}()

// 	// start worker
// 	for i := 0; i < worker; i++ {
// 		wg.Add(1)
// 		go workerJob(chINURL, quits[i], i)
// 	}

// 	go func() {
// 		for {
// 			if value := atomic.LoadUint32(&countResult); value == uint32(lenArrURL) {
// 				for i := 0; i < worker; i++ {
// 					quits[i] <- 1
// 				}
// 				return
// 			}
// 		}
// 	}()
// 	wg.Wait()
// }
