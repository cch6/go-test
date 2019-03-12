package main

import (
	"fmt"
	"runtime"
)

var NUM = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func main() {
	var cpuNum = runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	runtime.GOMAXPROCS(1)
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go print(i, ch)
	}
	for {
		msg := <-ch
		// println(msg)

		fmt.Println(msg)
		fmt.Println(NUM)
	}

	// ticker := time.NewTicker(5 * time.Second)
	// for _ = range ticker.C {
	// 	fmt.Println(time.Now())
	// 	// time.Sleep(1 * time.Second)
	// 	fmt.Println(NUM)
	// }

}

func add(i int) {

}

func print(i int, ch chan string) {
	for {
		NUM[i]++
		ch <- fmt.Sprintf("msg%d", i)
	}
}
