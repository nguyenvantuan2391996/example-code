package main

import (
	"fmt"

	"example-code/algorithm/search"
	"example-code/algorithm/sort"
)

func main() {
	nums := []int{3, 2, 5, 0, 1, 8, 7, 6, 9, 4}
	fmt.Printf("Before: %v", nums)
	sort.QuickSort(nums)
	fmt.Printf("After: %v", nums)

	fmt.Printf("Value 8 has number of index: %v", search.BinarySearch(nums, 8))
}
