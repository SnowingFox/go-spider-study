package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for n := range c {
		fmt.Printf("%c\n", n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(c)

	return c
}

func bufferedChannel() {
	c := make(chan int)
	go worker(c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	time.Sleep(time.Millisecond)
}

func closeChannel() {
	c := make(chan int)
	go worker(c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[0:1]
	b = append(b, 6, 7, 8, 9)
	fmt.Printf("a", a)
	fmt.Println("b", b)
	// bufferedChannel()
	closeChannel()
}
