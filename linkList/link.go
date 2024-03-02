package Linklist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

//TRAVERSE LINKED LIST AND PRINT ELEMENT

func (l LinkedList) TraverseAllNode() {
	current := l.head

	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}

//INSERTION BEGGINING OF THE NODE 

func (l *LinkedList)InesrtBeggingOFNode(val int){
	node1 := &Node{data: val}

	second := l.head
	node1.next = second
	l.head = node1
	

}

//INSERTING AT THE END OF THE NODE

func (l *LinkedList)InsertAtEndOfNOde(val int){
	node1 := &Node{data: val}

	if l.head == nil{
		l.head= node1
		l.length = 1
		return
	}
	
	current := l.head
	
	
	for current.next != nil{
		
		current =current.next
	}
	
	current.next = node1
	l.length ++
	fmt.Println(l.length)
}

//INSERT MIDDLE OF THE NODE 

func(l *LinkedList)AddMiddleOfNode(val int){

	node1 := &Node{data: val}
	
	if l.head == nil{
		l.head = node1
		return
	}
	  current := l.head

	pos := l.length / 2

	for i:= 1 ;i<pos;i ++{
		

		current =current.next
		
	}
	node1.next = current.next
	current.next =node1
	current = node1
}


//DELETE NODE 



