package linkedlist

import (
	"log"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Value: val}
}

type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Push(item T) {
	if l.head == nil {
		l.head = NewNode(item)
		return
	}
	head := l.head
	n := NewNode(item)
	n.Next = head
	l.head = n
}

func (l *LinkedList[T]) Pop() T {
	if l.isEmpty() {
		log.Println("Empty Linked-list")
		return *new(T)
	}
	popped := l.head.Value
	l.head = l.head.Next
	return popped
}

func (l *LinkedList[T]) Find(key T) *Node[T] {
	return nil
}

func (l *LinkedList[T]) isEmpty() bool {
	return l.head == nil
}
