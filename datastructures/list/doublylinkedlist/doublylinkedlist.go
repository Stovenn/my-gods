package doublylinkedlist

import (
	"fmt"
	"reflect"
)

type DoublyLinkedListNode[T any] struct {
	previous *DoublyLinkedListNode[T]
	next     *DoublyLinkedListNode[T]
	data     T
}

type DoublyLinkedList[T any] struct {
	head *DoublyLinkedListNode[T]
	tail *DoublyLinkedListNode[T]
	len  int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) AddHead(item T) {
	if l.head == nil {
		l.head = &DoublyLinkedListNode[T]{previous: nil, next: nil, data: item}
		return
	}
	head := l.head
	n := &DoublyLinkedListNode[T]{previous: nil, next: head, data: item}
	head.previous = nil
	l.head = n
	l.len++
}

func (l *DoublyLinkedList[T]) AddTail(item T) {
	if l.head == nil {
		n := &DoublyLinkedListNode[T]{previous: nil, next: nil, data: item}
		l.head = n
		l.tail = n
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	n := &DoublyLinkedListNode[T]{previous: currentNode, next: nil, data: item}
	currentNode.next = n
	l.tail = n
	l.len++
}

func (l *DoublyLinkedList[T]) Insert(item T, position int) {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) RemoveNode(key T) T {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Find(key T) *DoublyLinkedListNode[T] {
	if l.head == nil {
		return nil
	}
	currentNode := l.head
	for currentNode != nil {
		if reflect.DeepEqual(currentNode.data, key) {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

func (l *DoublyLinkedList[T]) Contains(key T) bool {
	return l.Find(key) != nil
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *DoublyLinkedList[T]) Print() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
