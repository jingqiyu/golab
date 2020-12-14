package dp

import (
	"fmt"
)

// 前缀和的一类问题

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	var prefixSum []int = make([]int, len(nums)+1)

	//另外一种写法
	/*for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}*/

	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	fmt.Println(prefixSum)
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.prefixSum[j+1] - this.prefixSum[i]
}
