package stack

type Stack struct {
	capacity int32

	size int32
	/**
	 * pointer to the tail of the list
	 */
	tail *node
}

const (
	DefaultCapacity = 8
)

type node struct {
	value int
	pre   *node
}

func CreateStack() *Stack {
	return &Stack{
		capacity: int32(DefaultCapacity),
		size:     0,
		tail:     nil,
	}
}

func (s *Stack) Push(v int) {

	if s.size == s.capacity {
		panic("exceed capacity")
	}
	node := node{
		value: v,
		pre:   nil,
	}
	node.pre = s.tail
	s.tail = &node
	s.size++
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		panic("exceed capacity")
	}
	tail := s.tail
	pre := tail.pre
	s.tail = pre
	s.size--
	return tail.value
}
