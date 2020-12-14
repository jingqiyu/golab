package huisu

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	l := len(nums)
	var rs [][]int
	var path []int
	var vis = make([]bool, l)
	var dfsPermuteUnique func(idx int)
	dfsPermuteUnique = func(idx int) {
		if idx == l {
			rs = append(rs, append([]int{}, path...)) // copy 1个 path 添加到rs中
			return
		}

		for i, v := range nums {
			//当前节点已经被访问过了， 2，前面的节点访问了，并且前面的节点和当前节点一样
			if vis[i] || i > 0 && vis[i-1] == true && v == nums[i-1] {
				continue
			}
			vis[i] = true          //当前节点标记为已使用
			path = append(path, v) //当前节点加入到路径中
			dfsPermuteUnique(idx + 1)
			path = path[:len(path)-1]
			vis[i] = false

		}

	}
	dfsPermuteUnique(0)
	//fmt.Println(rs)
	return rs
}
