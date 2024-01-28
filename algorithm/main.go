package main

import (
	"fmt"

	"example-code/algorithm/search"
	"example-code/algorithm/sort"
)

func main() {
	nums := []int{3, 2, 5, 0, 1, 8, 7, 6, 9, 4}
	fmt.Println(fmt.Sprintf("Before: %v", nums))
	sort.QuickSort(nums)
	fmt.Println(fmt.Sprintf("After: %v", nums))

	fmt.Println(fmt.Sprintf("Value 8 has number of index: %v", search.BinarySearch(nums, 8)))
}
