package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal   // 中断通道
	complete  chan error       // 报错通道
	timeout   <-chan time.Time // 超时通道
	tasks     []func(int)
}

var (
	ErrTimeout   = errors.New("received timeout ")   // 任务执行超时返回
	ErrInterrupt = errors.New("received interrupt ") // 收到操作系统信号时返回
)

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt) // 接收所有中断信号

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		// 执行 task
		task(id)
	}
	return nil
}

// 尝试从通道里取一下数据，判断是否有中断
func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt: // 收到一个中断信号
		signal.Stop(r.interrupt) // 停止接收后面的所有事件
		return true
	default:
		return false
	}
}
