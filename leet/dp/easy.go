package dp

//https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/solution/
/**
 * 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）角
 * 开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
 *
 *
 *
 * param: [][]int grid
 * return: int
 */
//解题关键 ， dp方程
// dp(i,j) = max( dp(i-1,j) , dp(i,j-1) ) + grid(i,j)
// 其中第一列和第一行比较特殊  只有单一通路

func maxValue(grid [][]int) int {

	maxFunc := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	if len(grid) == 0 {
		return 0
	}
	//修正第一列 只能通过上面到达，所以他的最大值就是当前值 + 上面的数值，，直接在原数组解题 节约空间
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}

	var i = 1
	var j = 1
	for i = 1; i < len(grid); i++ {
		for j = 1; j < len(grid[i]); j++ {
			grid[i][j] = maxFunc(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}

	lastJ := len(grid[len(grid)-1]) - 1
	lastI := len(grid) - 1
	return grid[lastI][lastJ] // 防止只有1行数据的时候 。 下面那段循环其实不会执行
}
