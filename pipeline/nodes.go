package pipeline

import (
	"sort"
)

// ArraySource 处理数据源
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	// out = sort.Ints(out)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// InMemSort sort in mem
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		// sort
		sort.Ints(a)
		// out
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// merge arr
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	// out = sort.Ints(out)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}
