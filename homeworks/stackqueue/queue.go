package stackqueue

type Queue[K any] struct {
	store []K
}

func NewQueue[K any]() *Queue[K] {
	return &Queue[K]{}
}

func (q *Queue[K]) Push(v K) {
	q.store = append(q.store, v)
}

func (q *Queue[K]) Pop() {
	length := len(q.store)
	if length == 0 {
		return
	}
	q.store = q.store[1:]
}

func (q *Queue[K]) Front() K {
	var k K
	length := len(q.store)
	if length == 0 {
		return k
	}
	return q.store[0]
}

func (q *Queue[K]) Size() int {
	return len(q.store)
}

func (q *Queue[K]) IsEmpty() bool {
	if len(q.store) == 0 {
		return true
	}
	return false
}
