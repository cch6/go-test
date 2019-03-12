package main

import (
	"fmt"

	"github.com/cch6/go-test/pipeline"
)

func main() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(4, 7, 2, 21, 4, 78, 9, 2)),
		pipeline.InMemSort(
			pipeline.ArraySource(42, 731, 12, 221, 4, 178, 39, 22)))
	for v := range p {
		fmt.Println(v)
	}
}
