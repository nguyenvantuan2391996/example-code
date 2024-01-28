package main

func partition(nums []int) ([]int, []int) {
	curA, curB := -1, -1
	pivot := nums[len(nums)-1]

	for _, value := range nums {
		curA++
		if value <= pivot {
			curB++
			if curA > curB {
				nums[curA], nums[curB] = nums[curB], nums[curA]
			}
		}
	}

	return nums[:curB], nums[curB:]
}

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	left, right := partition(nums)
	QuickSort(left)
	QuickSort(right)
}
