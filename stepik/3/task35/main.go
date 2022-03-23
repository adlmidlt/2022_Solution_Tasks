package main

func main() {

}
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {

	res := make(chan int)
	go func() {
		defer close(res)

		select {
		case val := <-firstChan:
			res <- val * val
		case val := <-secondChan:
			res <- val * 3
		case <-stopChan:
			return
		}
	}()
	return res
}
