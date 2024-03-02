package Linklist

import (
	"fmt"
)


//MERGE SORTED TWO ARRAY

func (l *LinkedList) MergeTwoSortArray(list1 *Node, list2 *Node) *Node {
    l1 := list1
    l2 := list2

    var mergedList, tail *Node
    for l1 != nil && l2 != nil {
        var smaller *Node
        if l1.Data <= l2.Data {
            smaller = l1
            l1 = l1.Next
        } else {
            smaller = l2
            l2 = l2.Next
        }
        if mergedList == nil {
            mergedList = smaller
            tail = smaller
        } else {
            tail.Next = smaller
            tail = tail.Next
        }
    }

    // Append remaining elements of either list
    if l1 != nil {
        if mergedList == nil {
            mergedList = l1
        } else {
            tail.Next = l1
        }
    }

    if l2 != nil {
        if mergedList == nil {
            mergedList = l2
        } else {
            tail.Next = l2
        }
    }

    return mergedList
}

func (l *LinkedList) NodePrint(n *Node) {

	current := n
	if n == nil {
		fmt.Println("there is no more node there ")
		return
	}

	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
}
