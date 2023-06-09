package link

import "fmt"

// 线性表,同时也是顺序表

type LinearList interface {
	InsertList(value int)             // 插入数据到线性表
	DeleteList(index int)             // 按元素下标位置删除元素
	GetLocationByValue(value int) int // 按值查找元素位置
	GetValueByLocation(index int) int // 按元素位置查找值
	PrintList()                       // 打印线性表
	DestroyList()                     // 销毁线性表
}

type LineList struct {
	data []int
	len  int // 是真实的长度比下标多一
	cap  int
}

func (l *LineList) InsertList(value int) {
	if l.len >= l.cap {
		fmt.Println("InsertList fail")
	}

	l.data[l.len] = value
	l.len++
}

func (l *LineList) DeleteList(index int) {
	if index > l.len-1 {
		fmt.Println("DeleteList fail")
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
	l.len--
}

func (l *LineList) GetLocationByValue(value int) int {
	for i, lLen := 0, l.len; i < lLen; i++ {
		if l.data[i] == value {
			return i
		}
	}
	return -1
}

func (l *LineList) GetValueByLocation(index int) int {
	for i, lLen := 0, l.len; i < lLen; i++ {
		if i == index {
			return l.data[i]
		}
	}
	return -1
}

func (l *LineList) PrintList() {
	if l.len == 0 {
		fmt.Println("list is nil")
	}
	for i, lLen := 0, l.len; i < lLen; i++ {
		fmt.Println("index:", i, ", value:", l.data[i])
	}
}

func (l *LineList) DestroyList() {
	l.data = make([]int, l.cap)
	l.len = 0
}

func NewLineList(cap int) *LineList {
	return &LineList{
		data: make([]int, cap),
		len:  0,
		cap:  cap,
	}
}
