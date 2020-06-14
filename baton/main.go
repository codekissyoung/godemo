package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)
	wg.Add(1)
	go Runner(baton)
	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {

	// 本 runner 拿到接力棒
	runner := <-baton

	go Runner(baton) // 下一棒的 Runner 开始等待

	// 本runner 开始奔跑
	fmt.Println("Runner ", runner, " runing with baton")
	time.Sleep(2 * time.Second)
	fmt.Println("Runner ", runner, " to the Line")

	// 如果是 第 4 棒，则比赛结束
	if runner == 4 {
		fmt.Println("Runner ", runner, " Finished , Race Over ")
		wg.Done()
		return
	}

	time.Sleep(time.Second)
	// 本 runner 与 下一 runner 开始交接棒
	newRunner := runner + 1
	fmt.Printf("Runner %d Exchange with Runner %d \n", runner, newRunner)
	baton <- newRunner
}
