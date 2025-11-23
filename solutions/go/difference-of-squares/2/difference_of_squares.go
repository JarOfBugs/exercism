package diffsquares

// SquareOfSum sum the given n natural number and perform the square on the sum result
// returns The square of sum as integer
func SquareOfSum(n int) int {
	var sum int    
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum * sum
}

// SumOfSquares Perform the square of the given n natural number and sums them
// returns The sum of squares as integer
func SumOfSquares(n int) int {
	var sum int
    for i := 1; i <= n; i++ {
        sum += i * i
    }
    return sum
}

// Difference perform the subtraction between the square of sum and the sum of squares of the given n natural number
// returns The difference as integer
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
