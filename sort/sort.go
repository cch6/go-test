package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 5, 2, 7, 8, 3}
	sort.Ints(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
