package huisu

import (
	"fmt"
)

/**

回溯任督二脉打通之法
一般回溯算法应用于求解过程类似一棵树， 而在从某个节点计算后之后，要弹出这个节点，继续计算其他节点的过程

常见问题； 全排列、组合和、  皇后等
常见末魔板

因为是树 ，我们认为解是一条从根节点到叶子节点的path  so .. 参数中需要1个解路径path = []int
因为要保存最终结果 需要一个 []path 数组  = [][]int
如果问题需要剪枝 那么一定需要一个开始位置 begin or start 命名随意
需要分叉的数组 一般是给定的输入  []int src
需要目标 规定达成的退出时机 target

一般模板如下：

func backtrack(result *[][]int, path []int, begin int , src []int, target int)  {
	if 一些情况已经可以返回的 {
		if target 达成 {
			result = append(result ,path )
		}
	}

	//分叉 走所有通路
	for i:=begin ;i < len(src) ;i ++ {
		path = append(path, src[i]) //把当前节点加入到路径中
		backtrack(result, path, i, src, 新的target) // 主要是目标更新，在新的子树中继续求解答案的过程
		path = path[:len(path)-1] //关键，这个元素用完了，就要删掉，回溯的过程，
	}

	return
}

*/

func allCombo(src []int) [][]int {
	var result [][]int
	var path []int
	backtrack(&result, path, src)
	return result
}

func inArray(a int, l []int) bool {
	for _, v := range l {
		if v == a {
			return true
		}
	}
	return false
}

func backtrack(result *[][]int, path []int, nums []int) {

	if len(path) == len(nums) {
		*result = append(*result, path)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 可以理解为剪枝  1个数字已经放到路径里面了，那么下面的分支都不需要再带着这个数字了
		if inArray(nums[i], path) {
			continue
		}
		path = append(path, nums[i])
		backtrack(result, path, nums)
		path = path[:len(path)-1] //关键，这个元素用完了，就要删掉，回溯的过程，
	}

}

// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
/**
 * 输入: k = 3, n = 7

输出: [[1,2,4]]
*/
func sumTarget(k int, n int) [][]int {
	var result [][]int
	var path []int
	backTrackSumTarget(&result, path, n, k, 1)
	return result
}

func backTrackSumTarget(result *[][]int, path []int, N int, k int, start int) {
	if len(path) == k || N <= 0 {
		if len(path) == k && N == 0 {
			*result = append(*result, path)
		}
		return

	}

	for i := start; i <= 9; i++ {
		if N < i {
			continue
		}
		path = append(path, i)
		backTrackSumTarget(result, path, N-i, k, i+1)
		path = path[:len(path)-1] //关键，这个元素用完了，就要删掉，回溯的过程，

	}

}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	backtrackCombinationSum(&result, path, candidates, 0, target)
	return result
}

func backtrackCombinationSum1(result *[][]int, path []int, candidates []int, start int, target int) {

	if len(path) == len(candidates) || target <= 0 {
		if target == 0 {
			*result = append(*result, path)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		backtrackCombinationSum1(result, path, candidates, i+1, target-candidates[i])
		//backtrack
		path = path[:len(path)-1]
	}

	return
}

func backtrackCombinationSum(result *[][]int, path []int, candidates []int, start int, target int) {

	if target <= 0 {
		if target == 0 {
			var dst []int
			dst = append(dst, path...)
			*result = append(*result, dst)
		}
		return
	}

	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		//这里做剪枝的时候一定要注意 ，元素可以重复利用 ，所以应该是 i  而不是 i+1
		backtrackCombinationSum(result, path, candidates, i, target-candidates[i])
		//backtrack
		path = path[:len(path)-1]
	}

	return
}

//输入：S = "qwe" 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
func permutation(S string) []string {
	var result [][]string
	var path []string
	var src []string
	for i := range S {
		src = append(src, string(S[i]))
	}
	backtrackPermutation(&result, path, src, 0)

	var ret []string
	for _, v := range result {
		var s string
		for _, v1 := range v {
			s += v1
		}
		ret = append(ret, s)
	}

	return ret
}

func backtrackPermutation(result *[][]string, path []string, src []string, begin int) {
	if len(path) == len(src) {
		var dst []string
		dst = append(dst, path...)
		*result = append(*result, dst)
		return
	}

	for i := 0; i < len(src); i++ {

		var flag bool
		for _, v := range path {
			if v == src[i] {
				flag = true
			}
		}
		if flag {
			continue
		}

		path = append(path, src[i])

		backtrackPermutation(result, path, src, i+1)

		path = path[:len(path)-1]

	}
	return
}

func generateParenthesis(n int) []string {
	var res []string

	dfsGenerateParenthesis(n, &res, 0, 0, "")
	fmt.Println(res)
	return res
}

func dfsGenerateParenthesis(n int, res *[]string, left int, right int, s string) {
	if left == n && right == n {
		*res = append(*res, s)
		return
	}
	if left < right {
		return
	}

	if left < n {
		dfsGenerateParenthesis(n, res, left+1, right, s+"(")
	}
	if right < n {
		dfsGenerateParenthesis(n, res, left, right+1, s+")")
	}
}
