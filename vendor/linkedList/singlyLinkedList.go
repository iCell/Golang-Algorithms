package linkedList

type SinglyNode struct {
	value interface{}
	next *SinglyNode
}

func NewSinglyNode(value interface{}) SinglyNode {
	return SinglyNode {
		value: value,
	}
}

type SinglyLinkedList struct {
	head *SinglyNode
}

func NewSinglyList(value interface{}) SinglyLinkedList {
	node := NewSinglyNode(value)
	return SinglyLinkedList {
		head: &node,
	}
}

func (l *SinglyLinkedList) Append(value interface{}) {
	node := NewSinglyNode(value)
	if l.head == nil {
		l.head = &node
	} else {
		l.head.next = &node
	}
}

func (l *SinglyLinkedList) Remove(value interface{}) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
	} else {
		last := l.head
		next := last.next
		for {
			if next == nil {
				break
			}
			if next.value == value {
				last.next = next.next
				break
			}
			last = next
			next = next.next
		}
	}
}

func (l *SinglyLinkedList) ToSlice() []interface{} {
	var results []interface{}
	node := l.head
	for {
		if node == nil {
			break
		}
		results = append(results, node.value)
		node = node.next
	}
	return results
}