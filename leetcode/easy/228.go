package easy

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var res []string
	start := nums[0]

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if start == nums[i] {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
			}

			if i+1 < len(nums) {
				start = nums[i+1]
			}
		}
	}

	return res
}
