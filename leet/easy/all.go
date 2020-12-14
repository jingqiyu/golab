package easy

import (
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	if len(arr) <= 2 {
		return true
	}
	sub := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {

		if arr[i]-arr[i-1] != sub {
			return false
		}
	}
	return true
}

func maximum69Number(num int) int {
	if num < 10 {
		return 9
	}
	var begin = 10

	if num > 1000 {
		begin = 1000
	} else if num > 100 {
		begin = 100
	} else {
		begin = 10
	}

	var div int
	var t int
	var prev int
	for num > 0 {
		div = num / begin
		if div == 9 {
			prev = 10*prev + div
		} else {
			t = num % begin
			return prev*begin*10 + 9*begin + t
		}
		num %= begin
		begin /= 10
	}
	return num
}

func maximum69Numbe2r(num int) int {
	/*sNum := strconv.Itoa(num)

	for i := range sNum {

		if string(sNum[i]) == "6" {
			sNum[i] = '9'
			break
		}
	}

	n, _ := strconv.Atoi(sNum)
	return n*/
	return 0
}

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	list := make([]int, 26)
	for _, value := range s {
		list[value-'a']++
	}
	for i := 0; i < len(s); i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
