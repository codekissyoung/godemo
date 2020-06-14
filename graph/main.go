package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var counter int
var mu sync.Mutex

func main() {

	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait()

	fmt.Println("Final Counter : ", counter)
}

func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mu.Lock() // ------------------ 临界区 ------------------
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mu.Unlock() // ----------------------------------------
	}
}

var shutdown int64

//
//func doWork(name string) {
//	defer wg.Done()
//	for {
//		fmt.Println("Doing ", name, " Work ")
//		time.Sleep(250 * time.Millisecond)
//		if atomic.LoadInt64(&shutdown) == 1 {
//			fmt.Println("Work ", name, " Shutting Down")
//			break
//		}
//	}
//}
