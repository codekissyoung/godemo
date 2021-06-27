package main

import (
	"log"
	"os"
	"time"

	"github.com/codekissyoung/godemo/runner"
)

const timeout = 10 * time.Second

func main() {
	log.Println("Start Working ")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("任务超时了")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("任务被中断")
			os.Exit(2)
		}
	}
	log.Println("Process ended ")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second * 2)
		log.Printf("------------- done ---------")
	}
}
