package recursion

func Factorial(n int) int {
	if n == 0 {

		return 1
	}

	return n * Factorial(n-1)
}

func SumOfDigits(n int) int {

	if n < 10 {
		return n
	}

	return n%10 + SumOfDigits(n/10)

}
