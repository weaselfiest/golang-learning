package easy

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		current := first + second
		first = second
		second = current
	}

	return second
}
