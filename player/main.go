package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 模拟乒乓球
func main() {
	court := make(chan int)

	wg.Add(1)
	go player("Link", court) // Link 准备接球
	court <- 1               // 裁判发球

	wg.Add(1)
	go player("Max", court) // Max 准备接 Link 打过来的球

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Println("Player ", name, " won")
			return
		}

		time.Sleep(time.Second / 2)
		n := rand.Intn(100)

		// 被 13 整除表示输球
		if n%13 == 0 {
			fmt.Println("Player ", name, " Missed")
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d \n", name, ball)

		ball++
		court <- ball
	}

}
