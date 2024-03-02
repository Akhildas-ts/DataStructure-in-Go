package Linklist

import "fmt"



//USING BINARY SERACH FOR SEARCH A ELEMENT FROM ARRAY ..

// frist of all we got staring point and end point, thats are left and right. 
// then we need to find the middle postion of the element 
//then  we are checking frist if the middle postion is the element 
// if middle index of value is greater than given value then we will check right side 
// or if middle index of value is less than then we will check the left side 


func BinarySerach(arr []int, x int) {

	left := 0
	right := len(arr) - 1

	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] == x {
			fmt.Println("value found at the index of :",mid,"value is :",x)
			return
		} 


		if arr[mid]<x{
			left = mid+1
		}else {
			right = mid -1
		}
	}

}

// GET THE LAST ELEMENT OF THE TARGET 

func BinarySerachOfLastElement(arr[]int,x int)int{

	left := 0
	right := len(arr)-1
	result := 0

	for left <= right{
		
		mid := left + (right - left)/2

		if arr[mid]==x{
			result = mid
			left = mid + 1
		}

		if arr[mid] <x{
			left =mid+1
		}else if arr[mid] > x{
			right = mid-1
		}
	}

	fmt.Println("these index ",result,"of ",x)
	return result
}