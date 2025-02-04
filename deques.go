package main 

import (
    "fmt"
    "iter"
)

/* 
A Deque, or a double-ended queue is ordered left to right, i.e. the first node
is the left-most node, and the last node is the right-most node 
*/
type Deque[T any] struct {
    first *Node[T]
    last *Node[T]
}

/*
The next node is the node adjacent to the right of this one, and the previous node
refers to the one to the left of this one
*/
type Node[T any] struct {
    prev *Node[T]
    next *Node[T]
    val T
}

func (deque Deque[T]) PrintDeque(flag string) {
    fmt.Println("Deque:", deque, "First:", deque.first, "Last:", deque.last)
    if flag == "forwards" {
        for n := range deque.TraverseForward() {
            fmt.Println("Node", n, ":", "val:", n.val, n.prev, ",", n.next)
        }
    } else {
        for n := range deque.TraverseBackwards() {
            fmt.Println("Node", n, ":", "val:", n.val, n.prev, ",", n.next)
        }
    }
}

func (deque Deque[T]) TraverseForward() iter.Seq[*Node[T]] {
    return func(yield func(*Node[T]) bool) {
        for e := deque.first; e != nil; e = e.next {
            if !yield(e) {
                return
            }
        }
    }
}

func (deque Deque[T]) TraverseBackwards() iter.Seq[*Node[T]] {
    return func(yield func(*Node[T]) bool) {
        for e := deque.last; e != nil; e = e.prev {
            if !yield(e) {
                return
            }
        }
    }
}

func (deque Deque[T]) NodeExists(node *Node[T]) bool {
    for n := range deque.TraverseForward() {
        if n == node {
            return true
        }
    }
    return false
}

func (deque *Deque[T]) InsertFirst(firstNode *Node[T]) {
    if deque.first == nil {
        deque.first = firstNode
        deque.last = firstNode
        firstNode.prev = nil
        firstNode.next = nil
    } else {
        deque.InsertBefore(deque.first, firstNode)
    }
}

func (deque *Deque[T]) InsertLast(firstNode *Node[T]) {
    if deque.last == nil {
        deque.first = firstNode
        deque.last = firstNode
        firstNode.prev = nil
        firstNode.next = nil
    } else {
        deque.InsertAfter(deque.last, firstNode)
    }
}

func (deque *Deque[T]) InsertAfter(node *Node[T], newNode *Node[T]) {
    if !deque.NodeExists(node) {
        return
    }
    newNode.prev = node
    if node.next == nil {
        deque.last = newNode
    } else {
        newNode.next = node.next
        node.next.prev = newNode
    }
    node.next = newNode
}

func (deque *Deque[T]) InsertBefore(node *Node[T], newNode *Node[T]) {
    if !deque.NodeExists(node) {
        return
    }
    newNode.next = node
    if node.prev == nil {
        deque.first = newNode 
    } else {
        newNode.prev = node.prev
        node.prev.next = newNode
    }
    node.prev = newNode
}

func (deque *Deque[T]) PopBack() {
    if deque.first != deque.last {
        deque.last = deque.last.prev
        deque.last.next = nil
    } else {
        // the front is the back because there was one element
        deque.first = nil
        deque.last = nil
    }
}

func (deque *Deque[T]) PopFront() {
    if deque.first != deque.last {
        deque.first = deque.first.next
        deque.first.prev = nil
    } else {
        // the front is the back because there was one element
        deque.first = nil
        deque.last = nil
    }
}

func (deque Deque[T]) Len() int {
    count := 0
    for range deque.TraverseForward() {
        count++
    }
    return count
}

