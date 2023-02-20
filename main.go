package main

import (
	"fmt"
	"github.com/stovenn/my-gods/datastructures/list/doublylinkedlist"
)

func main() {
	dl := doublylinkedlist.NewDoublyLinkedList[int]()
	dl.AddHead(30)
	fmt.Printf("%+v\n", dl)
	dl.AddHead(100)
	fmt.Printf("%+v\n", dl)
}
