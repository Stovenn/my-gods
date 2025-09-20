package main

import (
	"fmt"
	"my-gods/datastructures/heap"
)

func main() {
	h := new(heap.MinHeap)
	(*h) = []int{3, 5, 11, 8, 2, 4, 9, 32, 1}
	h.BuildHeap()
	fmt.Println(h)
}
