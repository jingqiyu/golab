package huisu

import (
	"fmt"
	"testing"
)

func Test_allCombo(t *testing.T) {
	combo := allCombo([]int{1, 2, 3})
	fmt.Println(combo)
}

func Test_sumTarget(t *testing.T) {
	fmt.Println(sumTarget(3, 9))
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func Test_permutation(t *testing.T) {
	permutation("Abc")
}

func Test_generateParenthesis(t *testing.T) {
	generateParenthesis(3)
}
