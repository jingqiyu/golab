package huisu

import (
	"fmt"
)

func Pailie(nums []int) [][]int {

	var result [][]int
	backtrackSwapWithFilter(&result, nums, 0)

	fmt.Println(result)
	return result
}

func backtrackSwap(result *[][]int, nums []int, k int) {
	if k == len(nums) {

		var tmp = make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	for i := k; i < len(nums); i++ {
		nums[i], nums[k] = nums[k], nums[i]
		backtrackSwap(result, nums, k+1)
		nums[i], nums[k] = nums[k], nums[i] // 重新换回来
	}
	return
}

func backtrackSwapWithFilter(result *[][]int, nums []int, k int) {

	if k == len(nums) {

		var tmp = make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	for i := k; i < len(nums); i++ {

		if i != k && nums[i] == nums[k] {
			backtrackSwapWithFilter(result, nums, k+1)
			return
			//continue
		}

		nums[i], nums[k] = nums[k], nums[i]
		backtrackSwapWithFilter(result, nums, k+1)
		nums[i], nums[k] = nums[k], nums[i] // 重新换回来
	}
	return
}
