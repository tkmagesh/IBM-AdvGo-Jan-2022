package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	mutex     *sync.Mutex
	closed    bool
}

/*
Fill in the blanks
*/

var ErrInvalidPoolSize = errors.New("Invalid pool size")
var ErrPoolClosed = errors.New("Pool closed")

func New(factory func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
		mutex:     &sync.Mutex{},
		closed:    false,
	}, nil
}

func (pool *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-pool.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire : From Pool")
		return r, nil
	default:
		fmt.Println("Acquire : From Factory")
		return pool.factory()
	}
}

func (pool *Pool) Release(r io.Closer) error {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.closed {
		r.Close()
		return ErrPoolClosed
	}

	select {
	case pool.resources <- r:
		fmt.Println("Release : In Pool")
		return nil
	default:
		fmt.Println("Release : Close & discard the resource")
		return r.Close()
	}
}

func (pool *Pool) Close() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.closed {
		return
	}
	pool.closed = true
	close(pool.resources)
	for r := range pool.resources {
		r.Close()
	}
}
