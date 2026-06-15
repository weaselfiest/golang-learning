package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	lastIndex := make(map[int]int)

	for i, num := range nums {
		if prev, ok := lastIndex[num]; ok {
			if i-prev <= k {
				return true
			}
		}

		lastIndex[num] = i
	}

	return false
}
