package main

import (
	"fmt"
	"os"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func tryDefer() {
	a := 1
	fmt.Println(1, a)
	a++
	defer fmt.Println(2, a)
	a++
	fmt.Println(3, a)
	a++
	defer fmt.Println(4, a)
	os.Create("12")
}

func main() {
	tryDefer()
	// fib := fibonacci()
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
	// fmt.Println(fib())
}
