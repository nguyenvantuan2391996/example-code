package search

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}

	return -1
}
