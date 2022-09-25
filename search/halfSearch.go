package search

import "fmt"

// 二分(折半) 查找, 返回下标

func HalfSearch(arr []int, target int) (int, bool) {
	if len(arr) <= 0 {
		fmt.Println("arr<=0, error")
		return 0, false
	}
	var left, right, mid int
	left = 0
	right = len(arr) - 1
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] == target {
			return mid, true
		}
		if arr[mid] < target {
			left = mid + 1
		}
		if arr[mid] > target {
			right = mid - 1
		}
	}

	return 0, false
}
