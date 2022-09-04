package search

import "fmt"

// 二分(折半) 查找, 返回下标

func HalfSearch(arr []int, target int) int {
	if len(arr) <= 0 {
		fmt.Println("arr<=0, error")
		return -1
	}
	var left, right, mid int
	left = 0
	right = len(arr) - 1
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		}
		if arr[mid] > target {
			right = mid - 1
		}
	}

	fmt.Println("target not found")
	return -1
}
