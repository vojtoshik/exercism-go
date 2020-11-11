package diffsquares

func SquareOfSum(n int) int {

	var s = 0

	for i:= 1; i <= n; i++ {
		s = s + i
	}

	return s * s
}

func SumOfSquares(n int) int  {
	var s = 0

	for i:= 1; i <= n; i++ {
		s = s + i * i
	}

	return s
}

func Difference(n int) int  {
	return SquareOfSum(n) - SumOfSquares(n)
}