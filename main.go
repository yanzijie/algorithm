package main

import (
	"fmt"
	"github.com/yanzijie/algorithm/search"
	"github.com/yanzijie/algorithm/sort"
)

func main() {
}

// TestHalfSearch 测试折半查找
func TestHalfSearch() {
	arr := []int{2, 3, 5, 7, 9, 12, 14, 16}
	target := 5
	index, ok := search.HalfSearch(arr, target)
	if !ok {
		fmt.Println("target:", target, "not exist")
	} else {
		fmt.Println("arr:", arr)
		fmt.Println("target:", target, "is index:", index)
	}
}

func TestSort() {
	arr := []int{33, 23, 12, 5, 8, 54, 24, 4}
	arr = sort.QuickSort(arr, 0, 7)
	fmt.Println("after QuickSort, arr:", arr)
}
