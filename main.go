package main

import (
	"fmt"
	"github.com/yanzijie/algorithm/link"
	"github.com/yanzijie/algorithm/search"
	"github.com/yanzijie/algorithm/sort"
)

func main() {
	TestOneWayLoopLink()
}

// TestOneWayLinkList 测试单向链表
func TestOneWayLinkList() {
	list := link.NewOneWayLinkList()

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

// TestOneWayLoopLink 测试单向循环链表
func TestOneWayLoopLink() {
	oneLoopLink := link.NewOneLoopLink()
	oneLoopLink.InsertToHead(4)
	oneLoopLink.InsertToHead(3)
	oneLoopLink.InsertToHead(2)
	oneLoopLink.InsertToHead(1)
	oneLoopLink.PrintfLink()

	oneLoopLink.InsertToTail(5)
	oneLoopLink.InsertToTail(6)
	oneLoopLink.InsertToTail(7)
	oneLoopLink.InsertToTail(8)
	oneLoopLink.PrintfLink()

	res := oneLoopLink.InsertByIndex(5, 18)
	if res {
		oneLoopLink.PrintfLink()
	}

	res = oneLoopLink.DeleteByIndex(5)
	if res {
		oneLoopLink.PrintfLink()
	}

}

// TestTwoWayLoopLinkList 测试双向循环链表
func TestTwoWayLoopLinkList() {
	list := link.NewDoubleList()

	list.InsertHead(4)
	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)
	list.PrintList()

	list.InsertTail(5)
	list.InsertTail(6)
	list.InsertTail(7)
	list.PrintList()

	list.InsertByIndex(16, 5)
	list.PrintList()

	list.DeleteByIndex(3)
	list.PrintList()

	list.DeleteByValue(6)
	list.PrintList()

	list.UpdateByIndex(1, 11)
	list.UpdateByIndex(6, 66)
	list.PrintList()

	list.UpdateByValue(11, 111)
	list.UpdateByValue(66, 666)
	list.UpdateByValue(2, 22)
	list.UpdateByValue(232, 444)
	list.PrintList()

	v1 := 666
	index, ok := list.GetIndexByValue(v1)
	if !ok {
		fmt.Println("GetIndexByValue error, value:", v1, "not exist")
	} else {
		fmt.Println("GetIndexByValue value:", v1, ",index:", index)
	}

	v2 := 4343
	index, ok = list.GetIndexByValue(v2)
	if !ok {
		fmt.Println("GetIndexByValue error, value:", v2, "not exist")
	} else {
		fmt.Println("GetIndexByValue value:", v2, ",index:", index)
	}

	v3 := 6
	value, ok := list.GetValueByIndex(v3)
	if !ok {
		fmt.Println("GetValueByIndex error, index:", v3, "invalid")
	} else {
		fmt.Println("GetValueByIndex index:", v3, ",value:", value)
	}

	v4 := 8
	value, ok = list.GetValueByIndex(v4)
	if !ok {
		fmt.Println("GetValueByIndex error, index:", v4, "invalid")
	} else {
		fmt.Println("GetValueByIndex index:", v4, ",value:", value)
	}

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
