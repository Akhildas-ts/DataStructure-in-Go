
// func (l *LinkedList) RemoveLastNelement(head *Node, n int) *Node {

// 	current := head
// 	length := 1

// 	for current.Next != nil {
// 		current = current.Next

// 		length++

// 	}
// 	current = head

// 	for i := 1; i <= length; i++ {
// 		fmt.Println("length is s", length)

// 		if n == length {
// 			second := head.Next
// 			head = second
// 			current = head

// 			// l.NodePrint(head)
// 			return head

// 		}

// 		if i == length-n {

// 			fmt.Println("current ", current.Data)
// 			fmt.Println("current.nex i == lengtht", current.Next.Data)
// 			// fmt.Println("current.next.next", current.Next.Next.Data)
// 			if current.Next != nil && current.Next.Next == nil{
// 				fmt.Println("im from sdfa")
// 				current.Next = nil
// 				break
// 			}

// 			current.Next = current.Next.Next

// 		} else {
// 			if current.Next == nil {
// 				break
// 			}

// 			fmt.Println("current data ",current.Data)
// 			fmt.Println("current next..",current.Next)
// 			current = current.Next
// 		}
// 	}
	 

// 	return current

// }

