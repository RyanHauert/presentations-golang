package main

import (
	"fmt"
	"time"
)

type Result struct {
	Worker string
	Job    int
}

type Worker struct {
	Name string
}

func (w Worker) DoWork(queue <-chan int, results chan<- *Result, complete chan<- bool) {
	for i := range queue {
		time.Sleep(100 * time.Millisecond)
		results <- &Result{w.Name, i}
	}
	complete <- true
}

type WorkerPool struct {
	workers []*Worker
	Queue   chan<- int
	Results <-chan *Result
}

func NewWorkerPool(workerCount int) *WorkerPool {
	queue := make(chan int, 100)
	results := make(chan *Result, 100)
	complete := make(chan bool)
	workers := make([]*Worker, workerCount)

	for i := 0; i < workerCount; i++ {
		worker := Worker{fmt.Sprintf("Worker%d", i)}
		go worker.DoWork(queue, results, complete)
		workers[i] = &worker
	}

	go func() {
		for i := 0; i < workerCount; i++ {
			<-complete
		}
		close(results)
	}()

	return &WorkerPool{workers, queue, results}
}
