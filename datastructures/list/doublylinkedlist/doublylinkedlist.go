package doublylinkedlist

type DoublyLinkedListNode[T any] struct {
	previous *DoublyLinkedListNode[T]
	next     *DoublyLinkedListNode[T]
	data     T
}

type DoublyLinkedList[T any] struct {
	head *DoublyLinkedListNode[T]
	tail *DoublyLinkedListNode[T]
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
}

func (l *DoublyLinkedList[T]) AddTail(item T) {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Insert(item T, position int) {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) RemoveNode(key T) T {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Find(key T) *DoublyLinkedListNode[T] {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Contains(key T) bool {
	panic("implement me")
}

func (l *DoublyLinkedList[T]) Empty() {
}

func (l *DoublyLinkedList[T]) IsEmpty() bool {
	panic("implement me")
}
