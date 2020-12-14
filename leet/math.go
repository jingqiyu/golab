package leet

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var f0 int64
	var f1 int64
	f0 = 1
	f1 = 2

	for n > 2 {
		t := f1
		f1 = f0 + f1
		f0 = t
		n--
	}

	return int(f1 % 1000000007)
}
