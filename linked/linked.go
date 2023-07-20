package main

import "errors"

// single linked list
type Node struct {
    value int 
    next *Node
}


type LinkedList struct {
    head *Node
}

func (l *LinkedList) length() int {
    count, runner :=  0, l.head 
    for runner != nil {
        count++
        runner = runner.next
    }
    return count 
}

func (l *LinkedList) get(index int) (int, error) {
    length := l.length()

    if index >= length {
        return 0, errors.New("Out of bounds")
    }

    current := l.head
    position := 0

    for position < index {
        current = current.next
        position++
    }
    return current.value, nil
}


func Test() *LinkedList {
    start := &Node{0, &Node{1, &Node{2, nil}}}

    return &LinkedList{head: start}
}


func main() {
    println(Test().length())
    println(Test().get(2))
}
