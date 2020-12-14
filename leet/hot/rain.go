package hot

//https://leetcode-cn.com/problems/container-with-most-water/
//盛水容器，双指针法
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l := 0
	r := len(height) - 1
	var ret int
	for l < r {
		minLr := height[l]
		lflag := true
		if height[r] <= height[l] {
			minLr = height[r]
			lflag = false
		}
		area := (r - l) * minLr
		if area >= ret {
			ret = area
		}
		if lflag {
			l++
		} else {
			r--
		}
	}
	return ret
}
