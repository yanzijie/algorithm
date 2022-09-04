package main

import (
	"fmt"
	"github.com/yanzijie/algorithm/link"
	"github.com/yanzijie/algorithm/search"
	"github.com/yanzijie/algorithm/sort"
)

func main() {
	TestSort()
}

func TestOneWayLinkList() {
	list := link.InitOneWayLinkList()

	list.InsertOnHead(1)
	list.InsertOnHead(2)
	list.InsertOnHead(3)
	list.InsertOnHead(4)
	list.PrintLink()

	list.InsertOnEnd(5)
	list.InsertOnEnd(6)
	list.InsertOnEnd(12)
	list.PrintLink()

	list.InsertOnIndex(33, 4)
	list.PrintLink()

	list.DelDataOnIndex(5)
	list.PrintLink()
	list.DelDataOnIndex(45)
	list.PrintLink()

	list.DelData(1)
	list.PrintLink()

	value := 33
	res := list.FindValue(value)
	if res {
		fmt.Println("list found value:", value)
	} else {
		fmt.Println("list not found value:", value)
	}

	value2 := 46
	list.UpdateValueByIndex(5, value2)
	list.PrintLink()
	list.UpdateValueByIndex(12, value2)
	list.PrintLink()
}

func TestHalfSearch() {
	arr := []int{2, 3, 5, 7, 9, 12, 14, 16}
	target := 19
	index := search.HalfSearch(arr, target)
	if index != -1 {
		fmt.Println("arr:", arr)
		fmt.Println("target:", target, "is index:", index)
	}
}

func TestSort() {
	arr := []int{33, 23, 12, 5, 8, 54, 24, 4}
	arr = sort.QuickSort(arr, 0, 7)
	fmt.Println("after QuickSort, arr:", arr)
}
