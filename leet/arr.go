package leet

import (
	"fmt"
)

func luckyNumbers(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	visited := make([][]int, rows)
	for i, _ := range visited {
		visited[i] = make([]int, cols)
	}

	var rowMin = func(row int) {
		m := matrix[row][0]
		c := 0
		for j := 0; j < cols; j++ {
			if matrix[row][j] < m {
				m = matrix[row][j]
				c = j
			}
		}
		for j := 0; j < cols; j++ {
			if j != c {
				visited[row][j] = 1
			}
		}
	}

	var colMax = func(col int) {
		max := matrix[0][col]
		r := 0
		for i := 0; i < rows; i++ {
			if matrix[i][col] > max {
				max = matrix[i][col]
				r = i
			}
		}
		for i := 0; i < rows; i++ {
			if i != r {
				visited[i][col] = 1
			}
		}

	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visited[i][j] == 1 {
				continue
			}
			rowMin(i)
			colMax(j)
		}
	}

	var ret []int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if visited[i][j] == 0 {
				ret = append(ret, matrix[i][j])
			}
		}

	}

	return ret

}

func fmtArrArr(src [][]int) {
	for _, v := range src {
		fmt.Println(v)
	}
}

/**
 * 给你一个大小为 rows x cols 的矩阵 mat，其中 mat[i][j] 是 0 或 1，请返回 矩阵 mat 中特殊位置的数目 。

特殊位置 定义：如果 mat[i][j] == 1 并且第 i 行和第 j 列中的所有其他元素均为 0（行和列的下标均 从 0 开始 ），则位置 (i, j) 被称为特殊位置。



示例 1：

输入：mat = [[1,0,0],
            [0,0,1],
            [1,0,0]]
输出：1
解释：(1,2) 是一个特殊位置，因为 mat[1][2] == 1 且所处的行和列上所有其他元素都是 0

*/
func numSpecial(mat [][]int) int {
	cols := len(mat[0])
	rows := len(mat)

	csCols := make([]int, cols)
	csRows := make([]int, rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				csCols[j]++
				csRows[i]++
			}
		}
	}

	fmt.Println(csCols)
	fmt.Println(csRows)

	var cnt int

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if csCols[j] == 1 && csRows[i] == 1 && mat[i][j] == 1 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
	return cnt
}
