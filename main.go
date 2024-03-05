package main

import String "linked/linkList/string"

func main() {

	// l1 := Linklist.LinkedList{}

	// A1.INSERT HEAD OF THE NODE

	// l1.InesrtBeggingOFNode(10)
	// l1.InesrtBeggingOFNode(20)
	// l1.InesrtBeggingOFNode(30)

	//A2.INSERT END OF THE NODE
	// l1.InsertAtEndOfNOde(10)
	// l1.InsertAtEndOfNOde(20)
	// l1.InsertAtEndOfNOde(30)

	//A3.
	// l1.AddMiddleOfNode(7)

	//A4
	// l1.DeleteNOde(10)
	// l1.TraverseAllNode()

	//A5 SEARCHING VALUE

	// node1, ind, err := l1.Serachvalue(20)

	// fmt.Println("search value ", node1, "index:", ind, err)

	//A6 LENGTH OF LIST
	// l1.LengthOfList()
	// l1.TraverseAllNode()

	//A7 REVERSER THE LIST

	//REVERSE THE LIST OF NODE

	// l1.TraverseAllNode()
	// println()

	// l1.ReverserList()

	//A8 MERGE TWO SORT ARRAY

	// list2 := &Linklist.Node{Data: 12, Next: &Linklist.Node{Data: 25, Next: &Linklist.Node{Data: 30}}}

	// store := l1.MergeTwoSortArray(list1, list2)
	// println()

	// l1.NodePrint(store)

	//A9 Binary serach
	// arr := []int{1, 3, 3, 3, 3, 5}
	// x :=3
	// Linklist.BinarySerach(arr,x)
	//  Linklist.BinarySerachOfLastElement(arr,x)

	// store := Linklist.BinarySerachOf(arr, x)

	// fmt.Println("these is slice of store ", store)

	//A10 REMVOVE THE NTH ELEMENT
	// list1 := &Linklist.Node{Data: 10, Next: &Linklist.Node{Data: 20, Next: &Linklist.Node{Data: 40, Next: &Linklist.Node{Data: 60,Next: &Linklist.Node{Data: 70}}}}}
	// 	list2 := &Linklist.Node{Data: 10}
	// 	// store := l1.RemoveLastNelement(list1,1)
	//   l1.Head = list2

	// 	// println()
	// 	swao:= l1.SwapPairs(list2)

	// 	l1.NodePrint(swao)

	// list1 := &Linklist.Node{Data: 1, Next: &Linklist.Node{Data: 5, Next: &Linklist.Node{Data: 4,Next: &Linklist.Node{Data: 6}}}}

	//   l1.RemvoeDuplicate(list1)
	// store := l1.SortList(list1)

	// l1.NodePrint(store)

	//B1 BINARY SEARCH

	// arr := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 3
	// store := Linklist.BinarySeracH(arr, target)
	// fmt.Println("stored value", store)

	//B2 STRING LAST WORD LENGTH\

  s:= "hey we are"
	String.LengthLastWord(s)

}
