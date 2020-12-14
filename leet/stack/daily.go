package stack

import (
	"fmt"
)

// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//
//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/daily-temperatures
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思考：
// 用栈巧妙的把所有没找到比他大的元素依次入栈
// 对于后面遍历到的元素，如果发现比栈顶（实际保存索引） 对应的元素大，， 那么栈顶的元素（索引） 的结果就找到了，就是当前索引 i - 栈顶索引  就是他们的距离
// 如果最后栈不空，那么栈里的所有元素相当于没有找到比他们还大的元素，对应位置至0

func dailyTemperatures(T []int) []int {

	var stack []int
	var stackTop int
	var empty = func() bool {
		return len(stack) == 0
	}
	var push = func(a int) {
		stack = append(stack, a)
		stackTop = a
	}
	var pop = func() (int, bool) {
		if len(stack) == 0 {
			return 0, false
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) != 0 {
			stackTop = stack[len(stack)-1]
		}

		return tmp, true
	}

	var rs = make([]int, len(T))
	push(0)

	var i = 1
	for ; i < len(T); i++ {
		for !empty() && T[i] > T[stackTop] {
			idx, _ := pop()
			rs[idx] = i - idx
		}
		push(i)
	}

	rs[i-1] = 0

	for !empty() {
		idx, _ := pop()
		rs[idx] = 0

	}

	return rs
}

/**
 * 给你一个数组 prices ，其中 prices[i] 是商店里第 i 件商品的价格。

商店里正在进行促销活动，如果你要买第 i 件商品，那么你可以得到与 prices[j] 相等的折扣，其中 j 是满足 j > i 且 prices[j] <= prices[i] 的 最小下标 ，如果没有满足条件的 j ，你将没有任何折扣。

请你返回一个数组，数组中第 i 个元素是折扣后你购买商品 i 最终需要支付的价格。



示例 1：

输入：prices = [8,4,6,2,3]
输出：[4,2,4,2,3]
解释：
商品 0 的价格为 price[0]=8 ，你将得到 prices[1]=4 的折扣，所以最终价格为 8 - 4 = 4 。
商品 1 的价格为 price[1]=4 ，你将得到 prices[3]=2 的折扣，所以最终价格为 4 - 2 = 2 。
商品 2 的价格为 price[2]=6 ，你将得到 prices[3]=2 的折扣，所以最终价格为 6 - 2 = 4 。
商品 3 和 4 都没有折扣。

*/
func finalPrices(prices []int) []int {
	var stack []int
	var stackTop int
	var empty = func() bool {
		return len(stack) == 0
	}
	var push = func(a int) {
		stack = append(stack, a)
		stackTop = a
	}
	var pop = func() (int, bool) {
		if len(stack) == 0 {
			return 0, false
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) != 0 {
			stackTop = stack[len(stack)-1]
		}

		return tmp, true
	}

	var ret = make([]int, len(prices))
	for i, v := range prices {
		if empty() {
			push(i)
			continue
		}

		for !empty() && v < prices[stackTop] {
			topIdx, _ := pop()
			ret[topIdx] = prices[topIdx] - v
		}

		push(i)
	}
	for !empty() {
		idx, _ := pop()
		ret[idx] = prices[idx]
	}
	return ret
}

// removeDuplicates
/**
 * 给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

输入："abbaca"
输出："ca"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 * param: string S
 * return: string
*/
func removeDuplicates(S string) string {
	var stack []string
	var stackTop string
	var empty = func() bool {
		return len(stack) == 0
	}
	var push = func(a string) {
		stack = append(stack, a)
		stackTop = a
	}
	var pop = func() (string, bool) {
		if len(stack) == 0 {
			return "", false
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) != 0 {
			stackTop = stack[len(stack)-1]
		}

		return tmp, true
	}

	if S == "" {
		return ""
	}

	for i := range S {
		if string(S[i]) != stackTop || empty() {
			push(string(S[i]))
		} else {
			pop()
		}
	}

	var rs string

	var f = func(str string) string {
		var result string
		strLen := len(str)
		for i := 0; i < strLen; i++ {
			result = result + fmt.Sprintf("%c", str[strLen-i-1])
		}
		return result
	}

	for !empty() {
		b, _ := pop()
		rs += string(b)
	}
	rs = f(rs)
	return rs
}
