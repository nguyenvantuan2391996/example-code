package main

import (
	"strings"
)

func equal(a, b int) bool {
	if a > b || a < b {
		return false
	}

	return true
}

func equalKiss(a, b int) bool {
	return a == b
}

func contains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}

	return false
}

func containsKiss(s, sub string) bool {
	return strings.Contains(s, sub)
}

func getMonth(month int) string {
	if month == 1 {
		return "Jan"
	} else if month == 2 {
		return "Feb"
	} else if month == 3 {
		return "Mar"
	}

	return ""
}

func getMonthKiss(month int) string {
	switch month {
	case 1:
		return "Jan"
	case 2:
		return "Feb"
	case 3:
		return "Mar"
	default:
		return ""
	}
}

func main() {
}
