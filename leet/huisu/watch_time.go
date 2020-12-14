package huisu

import (
	"fmt"
)

var simple = []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}

var res [][]int

func readBinaryWatch(num int) []string {
	backtrackReadBinaryWatch([]int{}, 0, num)
	var ret []string
	for _, path := range res {
		ret = append(ret, path2Time(path))
	}
	//fmt.Println(ret)

	return ret
}

func path2Time(path []int) string {

	h := 0
	m := 0
	for _, v := range path {
		if v > 0 {
			h += v
		} else {
			m -= v
		}
	}

	if m < 10 {
		return fmt.Sprintf("%d:0%d", h, m)
	} else {
		return fmt.Sprintf("%d:%d", h, m)
	}

}

func checkTime(path []int, i int) bool {
	tmp := 0
	for _, v := range path {
		if v < 0 {
			tmp -= v
		}

	}
	if tmp+i >= 60 {
		return false
	}
	return true
}

func backtrackReadBinaryWatch(path []int, begin int, num int) {
	if len(path) == num {
		var tmp = make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := begin; i < len(simple); i++ {
		if i <= 3 {
			path = append(path, simple[i])
		} else {
			//如果分钟已经超过60 这个元素就没法放进去
			if !checkTime(path, simple[i]) {
				continue
			}
			path = append(path, -simple[i]) //通过正负来取分下是小时还是分钟
		}
		backtrackReadBinaryWatch(path, i+1, num)
		// 回溯
		path = path[:len(path)-1]
	}

}
