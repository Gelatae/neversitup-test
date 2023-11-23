package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 4, 1, 1, 4}
	fmt.Println(FindOddInt(arr))
}

func FindOddInt(nums []int) int {
	// there will be only 1 odd int, that mean most of number will appear atlest 2 times
	// therefore it will be at most N/2 even int and 1 odd int
	// create map to store count int with size (N/2) + 1
	// predefined size help reduce cost when capacity is full
	exist := make(map[int]int, (len(nums)/2)+1)

	// loop count interger and store in map
	for _, num := range nums {
		exist[num]++
	}
	// loop over map to find odd int
	for key, val := range exist {
		if val%2 != 0 {
			// can return immediately because there will be only one answer
			return key
		}
	}
	return 0
}
