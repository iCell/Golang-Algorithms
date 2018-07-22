package stack

type Node struct {
	Value *interface{}
}

type Stack struct {
	Nodes []*interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	s.Nodes = append(s.Nodes, &value)
}

func (s *Stack) Pop() interface{} {
	lastIndex := len(s.Nodes) - 1
	last := *(s.Nodes[lastIndex])
	s.Nodes = s.Nodes[:lastIndex]
	return last
}
