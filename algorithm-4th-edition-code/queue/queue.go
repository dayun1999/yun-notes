package queue

type (
	Queue struct {
		start, end *node
		length     uint
	}
	node struct {
		value interface{}
		next  *node
	}
)

// create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// put an item on the end of queue
func (q *Queue) Enqueue(value interface{}) {
	n := &node{value: value, next: nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

// remove the front of the queue
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return n.value
}

// return the number if items in the queue
func (q *Queue) Len() uint {
	return q.length
}

// return the first item in the queue without removing it
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}
