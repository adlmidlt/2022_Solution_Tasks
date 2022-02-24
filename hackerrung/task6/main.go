package main

func flippingBits(n int64) int64 {
	// Write your code here
	var max uint32 = (1<<31-1)<<1 + 1
	var result = uint32(n) ^ max

	return int64(result)
}

func main() {}
