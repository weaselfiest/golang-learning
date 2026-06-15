package easy

func reverseBits(num int) int {
	result := 0

	for i := 0; i < 32; i++ {
		result <<= 1
		result |= num & 1
		num >>= 1
	}

	return result
}
