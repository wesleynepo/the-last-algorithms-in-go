package main

import "fmt"

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func newTree() *Tree {
	return &Tree{1, &Tree{2, &Tree{4, nil, nil}, &Tree{5, nil, nil}}, &Tree{3, &Tree{6, nil, nil}, &Tree{7, nil, nil}}}
}

func main() {
    println(compare(newTree(), newTree()))
}

// Binary tree pre order traversal
// 1 2 3 4 5 6 7

func bfs(tree *Tree, search int) bool {
    queue := make([]*Tree, 0);
    queue = append(queue, tree)

    for (len(queue) > 0) {
        next := queue[0]
        if next.value == search {
            return true
        }
        queue = queue[1:]
        if next.left != nil {
            queue = append(queue, next.left)
        }
        if next.right != nil {
            queue = append(queue, next.right)
        }
    }

    return false
}

func compare(a *Tree, b *Tree) bool {
    if (nil == a && nil == b) {
        return true
    }
    
    if (nil == a || nil == b) {
        return false
    }

    if (a.value != b.value) {
        return false
    }


    return compare(a.left, b.left) && compare(a.right, b.right)
}


func postOrder(tree *Tree) {
    if (nil == tree) {
        return
    }
    postOrder(tree.left)

    postOrder(tree.right)

    fmt.Println(tree.value)
}

func inOrder(tree *Tree) {
    if (nil == tree) {
        return
    }
        inOrder(tree.left)
    fmt.Println(tree.value)

        inOrder(tree.right)
}

func preOrder(tree *Tree) {
    if (nil == tree) {
        return
    }
    fmt.Println(tree.value)
        preOrder(tree.left)
        preOrder(tree.right)
}
