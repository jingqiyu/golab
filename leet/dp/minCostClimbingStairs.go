package dp

/**
 * 数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:

输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
 示例 2:

输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。

*/

// 考虑 f[i] 代表 i索引最小花费  那么 f[i] =  min(f[i+1],f[i+2]) + cost[i]
// 考虑从后往前来解决这个问题

func minCostClimbingStairs(cost []int) int {
	var minOne = func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	var f1, f2 = 0, 0
	for i := len(cost) - 1; i >= 0; i-- {
		f1, f2 = minOne(f1, f2)+cost[i], f1
	}
	return minOne(f1, f2)
}
