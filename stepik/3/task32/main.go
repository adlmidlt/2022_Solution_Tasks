package main

func main() {

}

func task2(c chan string, s string) {
	for i := 0; i < 5; i++ {
		c <- s + " "
	}
}

func removeDuplicates(inputStream, outputStream chan string) {
	defer close(outputStream)

	getStr := <-inputStream
	outputStream <- getStr

	for v := range inputStream {
		if getStr == v {
			continue
		}

		outputStream <- v
		getStr = v
	}
}
