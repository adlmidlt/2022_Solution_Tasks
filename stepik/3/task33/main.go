package main

func main() {
	done := make(chan struct{})

	go func(a chan struct{}) {
		work()
		close(a)
	}(done)
	<-done
}

func work() {

}
