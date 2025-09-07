package Task_2

import (
	"sync"
	"sync/atomic"
)

// Counter 计数器
type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

// Incr 递增计数器
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 获取当前计数器的值
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Reset 重置计数器
func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}

type AtomicCounter struct {
	value int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (c *AtomicCounter) Incr() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func (c *AtomicCounter) Reset() {
	atomic.StoreInt64(&c.value, 0)
}
