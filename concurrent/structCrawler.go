package main

import (
	"sync"
)

type Pool struct {
	wg                  sync.WaitGroup
	numberOfWorker      int
	lengthBufferChannel int
	chanInput           chan interface{}
	chanResult          chan interface{}
}
