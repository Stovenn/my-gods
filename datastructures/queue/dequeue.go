package queue

type Dequeue[T any] struct {
	buffer []T
	count  int
}

func (d *Dequeue[T]) isEmpty() bool {
	return d.count == 0
}

func (d *Dequeue[T]) addFirst(item T) {
	d.buffer = append([]T{item}, d.buffer...)
	d.count++
}

func (d *Dequeue[T]) addLast(item T) {
	d.buffer = append(d.buffer, item)
	d.count++
}

func (d *Dequeue[T]) removeFirst() T {
	item := d.buffer[0]
	d.buffer = append(d.buffer[1:])
	d.count--
	return item
}

func (d *Dequeue[T]) removeLast() T {
	item := d.buffer[d.count-1]
	d.buffer = append(d.buffer[:d.count-1])
	d.count--
	return item
}
