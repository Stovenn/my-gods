package queue

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue interface {
	Enqueue(string)
	Dequeue()
	IsEmpty()
}

func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{}
}

type LinkedQueue[T any] struct {
	head  *Node[T]
	queue *Node[T]
}

func (l *LinkedQueue[T]) Enqueue(s T) {
	oldQueue := l.queue
	l.queue = &Node[T]{value: s, next: nil}
	if l.IsEmpty() {
		l.head = l.queue
	} else {
		oldQueue.next = l.queue
	}
}

func (l *LinkedQueue[T]) Dequeue() T {
	item := l.head.value
	l.head = l.head.next
	if l.IsEmpty() {
		l.queue = nil
	}
	return item
}

func (l *LinkedQueue[T]) IsEmpty() bool {
	return l.head == nil
}
