package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := MinHeap{7, 1, 5, 2, 4, 3}
	heap.Init(&arr)
	length := len(arr)
	for i := 0; i < length; i++ {
		fmt.Println(heap.Pop(&arr))
	}
}
