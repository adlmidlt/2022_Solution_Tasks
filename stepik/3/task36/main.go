package main

import "sync"

func main() {

}

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)

		var sum int
		for {
			select {
			case val := <-arguments:
				sum += val
			case <-done:
				res <- sum
				return
			}
		}
	}()

	return res
}

func merge2Channels1(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	c := make(chan int)

	go func() {
		defer close(c)
		wg := new(sync.WaitGroup)
		mu := new(sync.Mutex)

		var arr1, arr2 []int
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				mu.Lock()
				defer wg.Done()
				arr1 = append(arr1, <-in1)
				mu.Unlock()
			}()
			wg.Add(1)
			go func() {
				mu.Lock()
				defer wg.Done()
				arr2 = append(arr2, <-in2)
				mu.Unlock()
			}()
		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- arr1[i] + arr2[i]
		}
	}()
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	cIndex := make(chan int)
	cValue := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- fn(x1)
			}()
			go func() {
				done <- fn(x2)
			}()

			go func(i int) {
				cIndex <- i
				cValue <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-cIndex] = <-cValue
		}

		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}
