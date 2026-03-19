package sieve

func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	isComposite := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !isComposite[i] {
			for j := i * i; j <= n; j += i {
				isComposite[j] = true
			}
		}
	}

	result := make([]int, 0, n/2)
	for i := 2; i <= n; i++ {
		if !isComposite[i] {
			result = append(result, i)
		}
	}
	return result
}
