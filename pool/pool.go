package pool

import (
	"errors"
	"sync"
)

var (
	ErrPoolClosed = errors.New("Pool has been closed ")
)

type Resource interface {
	Close() error
}

type ResourceCreator = func() (Resource, error)

type Pool struct {
	mutex    sync.Mutex
	resource chan Resource
	create   ResourceCreator
	closed   bool
}

func New(fn func() (Resource, error), size uint) (*Pool, error) {

	p := new(Pool)
	p.create = fn
	p.resource = make(chan Resource, size)

	return p, nil
}

// 关闭池子
func Close(p *Pool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true

	// 关闭 chan
	close(p.resource)

	// 从 chan 里依次取出资源，一个个关闭
	for r := range p.resource {
		_ = r.Close()
	}
}

func (p *Pool) Get() (Resource, error) {
	select {
	case r, ok := <-p.resource: // chan 里还有，读出一个
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default: // 新建一个
		return p.create()
	}
}

func (p *Pool) Put(r Resource) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// 池子以及关闭
	if p.closed {
		return r.Close()
	}

	select {
	case p.resource <- r: // 如果chan可写入，则放回去
		return nil
	default: // 如果不可写入，直接销毁资源
		return r.Close()
	}

}
