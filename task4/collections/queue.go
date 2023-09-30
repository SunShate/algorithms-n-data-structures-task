package collections

type QueueNode struct {
	Station int
	Time    int
}

/**
 * Queue implementation
 */
type Queue[T any] struct {
	bucket []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		bucket: []T{},
	}
}

func (q *Queue[T]) AddLast(input T) {
	q.bucket = append(q.bucket, input)
}

func (q *Queue[T]) PollFirst() (T, bool) {
	if q.IsEmpty() {
		var dummy T
		return dummy, false
	}
	value := q.bucket[0]
	var zero T
	q.bucket[0] = zero // Avoid memory leak
	q.bucket = q.bucket[1:]
	return value, true
}

func (q *Queue[T]) IsEmpty() bool {
	if len(q.bucket) < 1 {
		return false
	} else {
		return true
	}
}
