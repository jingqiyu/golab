package lab

import (
	"fmt"
	"sync"
)

func Test2Channel() {
	var (
		wg            = &sync.WaitGroup{}
		reachedChan   = make(chan int, 11)
		unReachedChan = make(chan int, 11)
		testArr       = []int{10, 20, 30, 40, 50, 6, 9, 22, 5, 7, 15}
	)
	for _, v := range testArr {
		wg.Add(1)
		go func(reachedChan chan<- int, unReachedChan chan<- int, v int) {
			defer wg.Done()
			if v%3 == 0 {
				fmt.Printf("v3===%d\n", v)
				reachedChan <- v
			} else if v%5 == 0 {
				fmt.Printf("v5===%d\n", v)
				unReachedChan <- v
			} else {
				return
			}

		}(reachedChan, unReachedChan, v)
	}
	wg.Wait()
	close(reachedChan)
	close(unReachedChan)

	fmt.Println("__________________rec_______________")
	for v := range reachedChan {
		fmt.Println(v)
	}
	fmt.Println("__________________rec_______________")
	for v := range unReachedChan {
		fmt.Println(v)
	}
}
