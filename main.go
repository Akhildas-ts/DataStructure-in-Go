package main

import (
	Linklist "linked/linkList"
)

func main() {

	l1 := Linklist.LinkedList{}

	// A1.INSERT HEAD OF THE NODE

	// l1.InesrtBeggingOFNode(10)
	// l1.InesrtBeggingOFNode(20)
	// l1.InesrtBeggingOFNode(30)

	//A2.INSERT END OF THE NODE
	l1.InsertAtEndOfNOde(10)
	l1.InsertAtEndOfNOde(20)
	l1.InsertAtEndOfNOde(30)

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
}
