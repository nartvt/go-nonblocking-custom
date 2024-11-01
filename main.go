package main

import (
	"fmt"
	"sync"
)

type CustomNonBlockingQueue struct {
	queue    []interface{}
	lock     sync.Mutex
	capacity int
}

func NewCustomNonBlockingQueue(cap int) *CustomNonBlockingQueue {
	return &CustomNonBlockingQueue{
		queue:    make([]interface{}, 0, cap),
		capacity: cap,
	}
}

func (r *CustomNonBlockingQueue) Length() int {
	return len(r.queue)
}

func (r *CustomNonBlockingQueue) Capacity() int {
	return r.capacity
}

func (r *CustomNonBlockingQueue) Publish(val interface{}) bool {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.Length() == r.Capacity() {
		return false
	}

	r.queue = append(r.queue, val)
	return true
}

func (r *CustomNonBlockingQueue) Subscribe() (interface{}, bool) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.Length() == 0 {
		return nil, false
	}

	val := r.queue[0]
	r.queue = r.queue[1:]
	return val, true
}

func main() {
	const cap = 10
	r := NewCustomNonBlockingQueue(cap)

	for i := 0; i < cap; i++ {
		if r.Publish(i) {
			fmt.Println("Publish ", i, " to queue")
		} else {
			fmt.Println("Queue is full")
		}
	}

	for {
		if val, ok := r.Subscribe(); ok {
			fmt.Println("Subscribe value: ", val)
		} else {
			fmt.Println("Empty queue")
			break
		}
	}
	fmt.Println("End")
}
