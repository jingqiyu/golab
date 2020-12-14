package dp

//https://leetcode-cn.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/solution/yuan-su-he-xiao-yu-deng-yu-yu-zhi-de-zheng-fang-2/

func matrixBlockSum(mat [][]int, K int) [][]int {

	//计算前缀arr  p  p[i][j] 代表 注意 p的下标比 mat 大 1  运算时要换算
	row := len(mat)
	col := len(mat[0])
	p := make([][]int, row+1)
	for i, _ := range p {
		p[i] = make([]int, col+1)
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			p[i][j] = p[i-1][j] + p[i][j-1] - p[i-1][j-1] + mat[i-1][j-1]
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {

			// 这里对于i,j 位置，他的矩形4个顶点，从上到下，从左到右 分别是
			//  [i-k,j-k] [i-k,j+k] [i+k,j-k] [i+k , j+k]

			//那么这个三角形的加和其实等于
			// p右下角  - p左下角 - p右上角 + p左上角  可以画图推导

			//上边界
			up := i - K - 1
			if up < -1 {
				up = -1
			}
			//下边界
			down := i + K //和他同行 不用+1
			if down > row-1 {
				down = row - 1
			}
			//左边界
			left := j - K - 1
			if left < -1 {
				left = -1
			}
			right := j + K
			if right > col-1 {
				right = col - 1
			}

			mat[i][j] = p[down+1][right+1] - p[down+1][left+1] - p[up+1][right+1] + p[up+1][left+1]
		}
	}

	//
	return mat
}

func prefixSumArr(mat [][]int) [][]int {
	col := len(mat)
	row := len(mat[0])
	p := make([][]int, row+1)
	for i, _ := range p {
		p[i] = make([]int, col+1)
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < col+1; j++ {
			p[i][j] = p[i-1][j] + p[i][j-1] - p[i-1][j-1] + mat[i-1][j-1]
		}
	}

	return nil
}
