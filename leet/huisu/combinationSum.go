package huisu

import (
	"sort"
)

/**
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]

*/
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var rs [][]int
	var backtrack func(begin int, path []int, target int)

	backtrack = func(begin int, path []int, target int) {
		if target <= 0 {
			if target == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				rs = append(rs, tmp)
			}
			return
		}

		for i := begin; i < len(candidates); i++ {
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			backtrack(i+1, path, target-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{}, target)
	return rs
}

func combinationSum3(k int, n int) [][]int {
	var rs [][]int
	var backtrack func(begin int, path []int, target int)
	var candidates = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	backtrack = func(begin int, path []int, target int) {
		if len(path) == k {
			if target == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				rs = append(rs, tmp)
			}
			return
		}

		for i := begin; i < len(candidates); i++ {
			if i > begin && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			backtrack(i+1, path, target-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{}, n)
	return rs

}
