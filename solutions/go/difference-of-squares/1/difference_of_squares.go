package diffsquares

import "math"

// SquareOfSum sum the given n natural number and perform the square on the sum result
// returns The square of sum as integer
func SquareOfSum(n int) int {
	sum := 0
    for current := range n {
        sum += current +1
    }
    return int(math.Pow(float64(sum), 2))
}

// SumOfSquares Perform the square of the given n natural number and sums them
// returns The sum of squares as integer
func SumOfSquares(n int) int {
	sum := 0
    for current := range n {
        sum += int(math.Pow(float64(current + 1), 2))
    }
    return sum
}

// Difference perform the subtraction between the square of sum and the sum of squares of the given n natural number
// returns The difference as integer
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
