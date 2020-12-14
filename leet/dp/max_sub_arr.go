package dp

func maxSubArray(nums []int) int {
	// dp[i] 标识到i这个元素为止 ，最大的子数组之和

	// 推到
	// dp[i] = dp[i-1] + a[i] // dp[i-1] > 0
	// dp[i] = a[i] // dp[i-1] <= 0 副作用

	var dp = make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(dp); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	var max = dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
