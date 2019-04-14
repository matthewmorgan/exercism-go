// Package diffsquares computes the difference of squares
package diffsquares

// method SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// method SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// method Difference returns the SquareOfSum less the SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
