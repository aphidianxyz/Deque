package deque 

import (
    "testing"
)

func TestEmptyNode(t *testing.T) {
    var emptyDeque Deque[int]
    emptyDequeLen := emptyDeque.Len()
    if emptyDequeLen != 0 {
        t.Fatalf(`Deque has a length: %d, expected length of 0.`, emptyDequeLen)
    }
}

func TestInsertingOneNodeFirst(t *testing.T) {
    var deque Deque[int]
    node := Node[int]{val:1}
    deque.InsertFirst(&node)

    dequeLen := deque.Len()
    if dequeLen != 1 {
        t.Fatalf(`Deque has length of %d, expected length of 1.`, dequeLen)
    }
    if deque.first != &node && deque.last != &node {
        t.Fatalf(`1st node, %v and last node, %v, 
            not equal to expected inserted node, %v`, deque.first, deque.last, &node)
    }
}

func TestPoppingBackInOneNodeDeque(t *testing.T) {
    var deque Deque[int]
    node := Node[int]{val:1}
    deque.InsertLast(&node)

    dequeLen := deque.Len()
    if dequeLen != 1 {
        t.Fatalf(`Deque has length of %d, expected length of 1.`, dequeLen)
    }
    if deque.first != &node && deque.last != &node {
        t.Fatalf(`1st node, %v and last node, %v, 
            not equal to expected inserted node, %v`, deque.first, deque.last, &node)
    }
}

func TestPoppingFrontInOneNodeDeque(t *testing.T) {
    var deque Deque[int]
    node := Node[int]{val:1}
    deque.InsertFirst(&node)

    deque.PopFront()
    dequeLen := deque.Len()

    if dequeLen != 0 {
        t.Fatalf(`Deque has a length: %d, expected length of 0.`, dequeLen)
    }
    if deque.first != nil && deque.last != nil {
        t.Fatalf(`1st (%v) and last (%v) nodes are not null`, deque.first, deque.last)
    }
}

func TestPoppingFirstAndLastNodeBack(t *testing.T) {
    var deque Deque[int]
    node := Node[int]{val:1}
    deque.InsertFirst(&node)

    deque.PopBack()
    dequeLen := deque.Len()

    if dequeLen != 0 {
        t.Fatalf(`Deque has a length: %d, expected length of 0.`, dequeLen)
    }
    if deque.first != nil && deque.last != nil {
        t.Fatalf(`1st (%v) and last (%v) nodes are not null`, deque.first, deque.last)
    }
}

func TestInsertBeforeANode(t *testing.T) {
    var deque Deque[int]
    node1 := Node[int]{val:1}
    node2 := Node[int]{val:2}
    node3 := Node[int]{val:3}
    deque.InsertFirst(&node1)
    deque.InsertLast(&node3)

    deque.InsertBefore(&node3, &node2)
    dequeLen := deque.Len()

    if dequeLen != 3 {
        t.Fatalf("Deque has length %d, expected length of 3", dequeLen)
    }
    if deque.first != &node1 {
        t.Fatalf("First element in deque is %v, expected %v", deque.first, &node1)
    }
    if deque.last != &node3 {
        t.Fatalf("First element in deque is %v, expected %v", deque.last, &node3)
    }
    if deque.first.next != &node2 || deque.last.prev != &node2 {
        t.Fatalf(`1st element has %v as next node, last element has %v,
            as previous node, expected %v in both cases`, deque.first.next, deque.last.prev, &node2)
    }
}

func TestInsertAfterANode(t *testing.T) {
    var deque Deque[int]
    node1 := Node[int]{val:1}
    node2 := Node[int]{val:2}
    node3 := Node[int]{val:3}
    deque.InsertFirst(&node1)
    deque.InsertLast(&node3)

    deque.InsertAfter(&node1, &node2)
    dequeLen := deque.Len()

    if dequeLen != 3 {
        t.Fatalf("Deque has length %d, expected length of 3", dequeLen)
    }
    if deque.first != &node1 {
        t.Fatalf("First element in deque is %v, expected %v", deque.first, &node1)
    }
    if deque.last != &node3 {
        t.Fatalf("First element in deque is %v, expected %v", deque.last, &node3)
    }
    if deque.first.next != &node2 || deque.last.prev != &node2 {
        t.Fatalf(`1st element has %v as next node, last element has %v,
            as previous node, expected %v in both cases`, deque.first.next, deque.last.prev, &node2)
    }
}

func TestInsertAfterCurrentLastNode(t *testing.T) {
    var deque Deque[int]
    node1 := Node[int]{val:1}
    node2 := Node[int]{val:2}
    node3 := Node[int]{val:3}
    deque.InsertFirst(&node1)
    deque.InsertLast(&node2)

    deque.InsertAfter(deque.last, &node3)
    dequeLen := deque.Len()

    if dequeLen != 3 {
        t.Fatalf("Deque has length %d, expected length of 3", dequeLen)
    }
    if deque.first != &node1 {
        t.Fatalf("First element in deque is %v, expected %v", deque.first, &node1)
    }
    if deque.last != &node3 {
        t.Fatalf("First element in deque is %v, expected %v", deque.last, &node3)
    }
    if deque.first.next != &node2 || deque.last.prev != &node2 {
        t.Fatalf(`1st element has %v as next node, last element has %v,
            as previous node, expected %v in both cases`, deque.first.next, deque.last.prev, &node2)
    }
}

func TestInsertBeforeCurrentLastNode(t *testing.T) {
    var deque Deque[int]
    node1 := Node[int]{val:1}
    node2 := Node[int]{val:2}
    node3 := Node[int]{val:3}
    deque.InsertFirst(&node2)
    deque.InsertLast(&node3)

    deque.InsertBefore(deque.first, &node1)
    dequeLen := deque.Len()

    if dequeLen != 3 {
        t.Fatalf("Deque has length %d, expected length of 3", dequeLen)
    }
    if deque.first != &node1 {
        t.Fatalf("First element in deque is %v, expected %v", deque.first, &node1)
    }
    if deque.last != &node3 {
        t.Fatalf("First element in deque is %v, expected %v", deque.last, &node3)
    }
    if deque.first.next != &node2 || deque.last.prev != &node2 {
        t.Fatalf(`1st element has %v as next node, last element has %v,
            as previous node, expected %v in both cases`, deque.first.next, deque.last.prev, &node2)
    }
}

func TestGetNodeAtIndex(t *testing.T) {
    var deque Deque[int]
    node1 := Node[int]{val:1}
    node2 := Node[int]{val:2}
    node3 := Node[int]{val:3}
    node4 := Node[int]{val:4}
    node5 := Node[int]{val:5}
    deque.InsertLast(&node1)
    deque.InsertLast(&node2)
    deque.InsertLast(&node3)
    deque.InsertLast(&node4)
    deque.InsertLast(&node5)

    want := deque.At(3)
    if want != node3 {
        t.Fatalf(`Expected %v, but got %v at index 3 of deque`, node3, want)
    }
}
