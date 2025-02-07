package stackqueue

type Stack[K any] struct {
	store []K
}

func NewStack[K any]() *Stack[K] {
	return &Stack[K]{}
}

func (s *Stack[K]) Push(v K) {
	s.store = append(s.store, v)
}

func (s *Stack[K]) Pop() {
	length := len(s.store)
	if length == 0 {
		return
	}
	s.store = s.store[:len(s.store)-1]
}

func (s *Stack[K]) Top() K {
	var k K
	length := len(s.store)
	if length == 0 {
		return k
	}
	return s.store[len(s.store)-1]
}

func (s *Stack[K]) Size() int {
	return len(s.store)
}

func (s *Stack[K]) IsEmpty() bool {
	if len(s.store) == 0 {
		return true
	}
	return false
}
