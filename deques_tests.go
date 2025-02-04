package main 

import "fmt"

func main() {
    // -=-=-=- Tests -=-=-=-
    var deque Deque[int]
    node0 := Node[int]{val: 0}

    fmt.Println(deque.Len())
    deque.InsertFirst(&node0)
    deque.PrintDeque("forwards")
    fmt.Println(deque.Len())

    node1 := Node[int]{val: 1}
    deque.InsertAfter(&node0, &node1)
    deque.PrintDeque("forwards")

    nodeBetween01 := Node[int]{val:-1}
    deque.InsertBefore(&node1, &nodeBetween01)
    deque.PrintDeque("forwards")

    node2 := Node[int]{val: 2}
    node3 := Node[int]{val: 3}
    deque.InsertAfter(&nodeBetween01, &node2)
    deque.InsertBefore(&node2, &node3)
    deque.PrintDeque("forwards")

    deque.PopBack()
    deque.PrintDeque("forwards")

    deque.PopFront()
    deque.PrintDeque("forwards")

    /*
    var dequeBefore Deque[int]
    nodeBefore := node2
    dequeBefore.InsertLast(&nodeBefore)
    dequeBefore.PrintDeque("forwards")

    deque.PrintDeque("forwards")
    dequeBefore.PopFront()
    dequeBefore.PrintDeque("forwards")
    */

}

