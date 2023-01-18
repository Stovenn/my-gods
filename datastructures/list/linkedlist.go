package list

import (
	"log"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

func newNode[T any](val T) *Node[T] {
	return &Node[T]{value: val}
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Push(item T) {
	if l.head == nil {
		l.head = newNode(item)
		return
	}
	head := l.head
	n := newNode(item)
	n.next = head
	l.head = n
}

func (l *LinkedList[T]) Pop() T {
	if l.isEmpty() {
		log.Println("Empty Linked-list")
		return *new(T)
	}
	popped := l.head.value
	l.head = l.head.next
	return popped
}

func (l *LinkedList[T]) Find(key T) *Node[T] {
	return nil
}

func (l *LinkedList[T]) isEmpty() bool {
	return l.head == nil
}
