package main

import "fmt"

func main() {
	nums := []int{3, 2, 5, 0, 1, 8, 7, 6, 9, 4}
	fmt.Println(fmt.Sprintf("Before: %v", nums))
	QuickSort(nums)
	fmt.Println(fmt.Sprintf("After: %v", nums))
}
