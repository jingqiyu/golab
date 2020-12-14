package binary_search

func missingNumber(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i < j {
		mixIdx := (i + j) / 2
		if nums[mixIdx] > mixIdx {
			j = mixIdx - 1
		} else {
			i = mixIdx + 1
		}
	}
	if nums[i] == i {
		return i + 1
	}
	return i
}
