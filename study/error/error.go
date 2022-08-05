package main

import "fmt"

func somethingBoom() {
	panic("Boom!!!")
}

func someSafeFunc() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panic: %s", e)
		}
	}()

	somethingBoom()
	fmt.Println("After func")
}

func main() {
	someSafeFunc()
}
