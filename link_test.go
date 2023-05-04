package main

import (
	"fmt"
	"github.com/yanzijie/algorithm/link"
	"testing"
)

// 链表测试

// TestOneWayLinkList 测试单向链表
func TestOneWayLinkList(t *testing.T) {
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
func TestOneWayLoopLink(t *testing.T) {
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
	oneLoopLink.InsertToTail(9)
	oneLoopLink.PrintfLink()

	ok := oneLoopLink.InsertByIndex(3, 18)
	if ok {
		oneLoopLink.PrintfLink()
	}

	ok = oneLoopLink.DeleteByIndex(5)
	if ok {
		oneLoopLink.PrintfLink()
	}

	ok = oneLoopLink.DeleteByValue(7)
	if ok {
		oneLoopLink.PrintfLink()
	}

	ok = oneLoopLink.ModifyByIndex(2, 9981)
	if ok {
		oneLoopLink.PrintfLink()
	}

	ok = oneLoopLink.ModifyByValue(18, 81)
	if ok {
		oneLoopLink.PrintfLink()
	}

	oneLoopLink.SearchByIndex(6)
	oneLoopLink.SearchByValue(8)
}

// TestTwoWayLoopLinkList 测试双向循环链表
func TestTwoWayLoopLinkList(t *testing.T) {
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

func TestLinerList(t *testing.T) {
	l := link.NewLineList(10)
	l.InsertList(1)
	l.InsertList(2)
	l.InsertList(3)
	l.InsertList(4)
	l.InsertList(5)
	l.InsertList(6)
	l.InsertList(7)
	l.InsertList(8)
	l.PrintList()
	fmt.Println("after delete")
	l.DeleteList(3)
	l.PrintList()
	res := l.GetLocationByValue(6)
	if res == -1 {
		fmt.Println("GetLocationByValue error")
	} else {
		fmt.Println("index for value 6 is:", res)
	}

	res = l.GetValueByLocation(3)
	if res == -1 {
		fmt.Println("GetValueByLocation error")
	} else {
		fmt.Println("value for index 3 is:", res)
	}
	fmt.Println("destroy list")
	l.DestroyList()
	l.PrintList()
}
