package main

import (
	"fmt"
	"strings"
)

type node struct {
	value interface{}
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Stack struct {
	head *node
	len  int
}

func (s *Stack) push(value interface{}) {
	s.len += 1
	newNode := new(node)
	newNode.value = value
	if s.head == nil {
		s.head = newNode
	} else {
		current := s.head
		for ; current.next != nil; current = current.next {
		}
		current.next = newNode
	}
}

func (s *Stack) pop() string {
	s.len -= 1
	if s.head == nil {
		return "Nothing to pop out: the stack is empty"
	} else {
		if s.head.next == nil {
			return fmt.Sprintf("%v", s.head)
		}
		var prev *node
		current := s.head
		for ; current.next != nil; current = current.next {
			prev = current
		}
		elem := fmt.Sprintf("%v", current)
		prev.next = nil
		return elem
	}
}

func (s Stack) String() string {
	var sb strings.Builder
	sb.WriteString("STACK HEAD\n")
	for current := s.head; current != nil; current = current.next {
		sb.WriteString(fmt.Sprintf("%v\n", current))
	}
	sb.WriteString("STACK END")
	return sb.String()
}

func main() {
	s := new(Stack)
	for i := 0; i < 10; i++ {
		s.push(i)
	}
	fmt.Println(s)
	for i := 0; i < 10; i++ {
		fmt.Println(s.pop())
	}
}
