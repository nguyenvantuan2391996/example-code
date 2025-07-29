package utils

import "strings"

func Encode(s, secret string, num int) string {
	length := len(s)
	if length < num {
		return s
	}

	mid := length / 2
	start := mid - num/2
	end := mid + num/2

	if start < 0 {
		start = 0
	}

	if end >= length {
		end = length
	}

	arr := strings.Split(s, "")
	for i := start; i < end; i++ {
		arr[i] = secret
	}

	return strings.Join(arr, "")
}
