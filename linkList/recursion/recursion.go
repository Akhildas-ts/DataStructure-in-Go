package recursion

// FACTORIAL

func Factorial(n int) int {
	if n == 0 {

		return 1
	}

	return n * Factorial(n-1)
}

// SUM OF THE DIGITS

func SumOfDigits(n int) int {

	if n < 10 {
		return n
	}

	return n%10 + SumOfDigits(n/10)

}

//BINARY SEACH WITH RECURSION

func BinaryWithRecurion(arr []int, left int, right int, target int) int {

	if left <= right {

		if arr[left] == target {
			return left
		}

		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			return BinaryWithRecurion(arr, mid+1, right, target)
		} else {
			return BinaryWithRecurion(arr, mid-1, right, target)
		}
	}
return -1
}
