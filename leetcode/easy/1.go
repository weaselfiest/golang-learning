package easy

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{i, j}
		}
		seen[num] = i
	}
	panic("no solution")
}
