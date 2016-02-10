package cbuf

//CircularBuffer is a
type CircularBuffer struct {
	head   int
	tail   int
	size   int
	buffer []interface{}
}

//NewCircularBuffer is a constructor which returns a pointer to a CircularBuffer instance
func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{head: 0, tail: 0, size: size, buffer: make([]interface{}, size)}
}

//Enqueue an element to the CircularBuffer
func (cbuf *CircularBuffer) Enqueue(item interface{}) {
	cbuf.buffer[cbuf.head] = item
	cbuf.head++
	if cbuf.head == cbuf.size {
		cbuf.head = 0
	}
}

//Dequeue removes an item from the head of the CircularBuffer
func (cbuf *CircularBuffer) Dequeue() interface{} {
	item := cbuf.buffer[cbuf.tail]
	cbuf.tail++
	if cbuf.tail == cbuf.size {
		cbuf.tail = 0
	}
	return item
}

//Peak returns an element from the CircularBuffer without removing its
func (cbuf *CircularBuffer) Peak() interface{} {
	return cbuf.buffer[cbuf.head]
}
