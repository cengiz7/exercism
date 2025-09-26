package diffsquares

func SquareOfSums(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}