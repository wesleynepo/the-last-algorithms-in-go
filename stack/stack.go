package main 

import "errors"



type Node struct {
    value int
    prev *Node
}


type Stack struct {
    length int
    head *Node
}

func (s *Stack) push(value int) {
    node := &Node{value, nil}
    node.prev = s.head
    s.head = node
    s.length++
}

func (s *Stack) pop() (int, error) {
    if s.head != nil {
        s.length--
        value := s.head.value
        s.head = s.head.prev
        return value, nil
    }
    return 0, errors.New("Empty stack")
}

func (s *Stack) peek() (int, error) {
    if s.head != nil {
        return s.head.value, nil
    }

    return 0, errors.New("Empty stack")
}

func (s *Stack) print() {
    current := s.head
    for current != nil {
        println(current.value)
        current = current.prev
    }
}

func main() {
    stack := &Stack{0, nil}
    stack.push(1)
    stack.push(2)
    stack.print()
    stack.pop()
    println(stack.pop())
    stack.print()
}
