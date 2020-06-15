package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 20
)

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	taskChan := make(chan string, taskLoad)

	for workerID := 0; workerID < numberGoroutines; workerID++ {
		wg.Add(1)
		go worker(taskChan, workerID)
	}

	for post := 1; post <= taskLoad; post++ {
		taskChan <- fmt.Sprintf("Task : %d", post)
	}

	close(taskChan)

	wg.Wait()
}

func worker(taskChan chan string, worker int) {
	defer wg.Done()
	for {
		task, ok := <-taskChan
		if !ok {
			fmt.Println("worker ", worker, " shutding down")
			return
		}
		fmt.Println("worker ", worker, " start task ", task)
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		fmt.Println("worker ", worker, " Complete task ", task)
	}
}
