package linkedList

type DoubleNode struct {
	value interface{}
	last *DoubleNode
	next *DoubleNode
}

func NewDoubleNode(value interface{}) DoubleNode {
	return DoubleNode {
		value: value,
	}
}

type DoubleLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
}

func NewDoubleList(value interface{}) DoubleLinkedList {
	node := NewDoubleNode(value)
	return DoubleLinkedList {
		head: &node,
		tail: &node,
	}
}

func (l *DoubleLinkedList) Append(value interface{}) {
	node := NewDoubleNode(value)
	if l.head == nil {
		l.head = &node
		l.tail = &node
	} else if l.head == l.tail {
		l.head.next = &node
		node.last = l.head
		l.tail = &node
	} else {
		l.tail.next = &node
		node.last = l.tail
	}
}

func (l *DoubleLinkedList) Remove(value interface{}) {
	if l.head == nil {
		return
	} else if l.head.value == value {
		next := l.head.next
		next.last = nil
		l.head = next
	} else {
		node := l.head
		for {
			if node.value == value {
				last := node.last
				next := node.next
				if next != nil {
					next.last = last
				}
				if last != nil {
					last.next = next
				}
			}
			node = node.next
			if node == nil {
				break
			}
		}
	}
}

func (l *DoubleLinkedList) ToSlice() []interface{} {
	var result []interface{}
	node := l.head
	for {
		result = append(result, node.value)
		node = node.next
		if node == nil {
			break
		}
	}
	return result
}