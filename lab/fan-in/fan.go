package fan_in

import (
	"sync"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

type squareResult struct {
	src int
	dst int
}

func square(inCh <-chan int) <-chan squareResult {
	out := make(chan squareResult)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- squareResult{src: n, dst: n * n}
		}
	}()

	return out
}

func merge(cs ...<-chan squareResult) <-chan squareResult {
	out := make(chan squareResult)
	//	out := make(chan squareResult, 100)

	var wg sync.WaitGroup

	collect := func(in <-chan squareResult) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
