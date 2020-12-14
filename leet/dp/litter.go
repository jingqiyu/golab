package dp

//给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
//
//换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
//
//以数组形式返回答案。
// 思考 ，前缀和比较 ，因为元素都在0，100 建立一个101长度的数组， 标识nums[i] 的统计数据
// 在求得该数组的前缀和 那么 p[i] 的元素 就代表 比 i = nums[x] 小的元素的个数
func smallerNumbersThanCurrent(nums []int) []int {
	var p = make([]int, 101) // 考虑有 0  和 100 两个元素 那么 0 元素
	for i := 0; i < len(nums); i++ {
		p[nums[i]]++
	}
	p[0] = 0
	for i := 1; i < len(p); i++ {
		p[i] += p[i-1]
	}
	var rs = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		rs[i] = p[nums[i]-1]
	}
	return rs
}
