package Linklist

import (
	"fmt"
)
// )
// func (l *LinkedList) MergeTwoSortArray(list1 *Node, list2 *Node) *Node {
//     var mergedHead, tail *Node
//     l1 := list1
//     l2 := list2

//     for l1 != nil && l2 != nil {
//         var smallerNode *Node
//         if l1.Data <= l2.Data {
//             smallerNode = l1
//             l1 = l1.Next
//         } else {
//             smallerNode = l2
//             l2 = l2.Next
//         }

//         // If mergedHead is nil, assign the smaller node as the head
//         if mergedHead == nil {
//             mergedHead = smallerNode
//             tail = smallerNode
//         } else {
//             // Append the smaller node to the merged list
//             tail.Next = smallerNode
//             // Update the tail to the newly added node
//             tail = smallerNode
//         }
//     }

//     // Append the remaining nodes of list1 or list2
//     if l1 != nil {
//         if mergedHead == nil {
//             mergedHead = l1
//         } else {
//             tail.Next = l1
//         }
//     }

//     if l2 != nil {
//         if mergedHead == nil {
//             mergedHead = l2
//         } else {
//             tail.Next = l2
//         }
//     }

//     return mergedHead
// }


func (l *LinkedList) NodePrint(n *Node) {

	current := n
	if n == nil {
		fmt.Println("there is no more node there ")
		return
	}

	for current != nil {
		fmt.Print(current.Data," ")
		current = current.Next
	}
}
