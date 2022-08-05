package main

import (
	"fmt"
	"sync"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("Safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) read() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var num atomicInt
	num.increment()
	num.increment()
	fmt.Println(num.read())
}
