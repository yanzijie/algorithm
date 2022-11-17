package link

import (
	"fmt"
)

// 单向链表

type IOneWayLinkList interface {
	InsertOnHead(value int)                  //从链表头部插入
	InsertOnEnd(value int)                   //从尾部添加元素
	InsertOnIndex(value int, index int)      //在指定位置添加元素  index是下标，从0开始,
	DelDataOnIndex(index int)                //删除链表指定位置的节点, index是下标，从0开始
	DelData(value int)                       //删除链表指定值的节点
	PrintLink()                              //打印链表所有节点
	FindValue(value int) bool                //查询链表中是否包含某个值
	UpdateValueByIndex(index int, value int) //根据index修改值为value, index是下标，从0开始
}

// Node 节点
type Node struct {
	data int // 节点数据
	next *Node
}

// List 链表
type List struct {
	head   *Node  // 头节点
	length uint64 // 链表长度
}

func NewOneWayLinkList() IOneWayLinkList {
	return &List{
		head:   nil,
		length: 0,
	}
}

func (l *List) InsertOnHead(value int) {
	n := &Node{
		data: value,
		next: nil,
	}

	if l.head == nil {
		l.head = n
		l.length++
		fmt.Println("InsertOnHead success, this head")
		return
	}
	n.next = l.head
	l.head = n
	l.length++
	fmt.Println("InsertOnHead success")
}

func (l *List) InsertOnEnd(value int) {
	n := &Node{
		data: value,
		next: nil,
	}

	if l.head == nil {
		l.head = n
		l.length++
		fmt.Println("InsertOnEnd success, this head")
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	// 找到链表尾部了
	cur.next = n
	l.length++
	fmt.Println("InsertOnEnd success")
}

func (l *List) InsertOnIndex(value int, index int) {
	if index < 0 {
		fmt.Println("InsertOnIndex error, index error")
		return
	}

	n := &Node{
		data: value,
		next: nil,
	}

	if l.head == nil {
		l.head = n
		l.length++
		fmt.Println("InsertOnIndex success")
		return
	}

	if index == 0 {
		l.InsertOnHead(value)
	}

	i := 1
	cur := l.head
	for cur.next.next != nil {
		if i == index {
			n.next = cur.next
			cur.next = n
			l.length++
			fmt.Println("InsertOnIndex success")
			return
		}
		i++
		cur = cur.next
	}
	// 走到这里 cur.next.next 为nil, 说明可能要插在链表尾部的前一个, cur.next是最后一个元素
	if i == index {
		n.next = cur.next
		cur.next = n
		l.length++
		fmt.Println("InsertOnIndex success")
	} else {
		fmt.Println("InsertOnIndex error, index error")
	}
}

func (l *List) DelDataOnIndex(index int) {
	if index < 0 {
		fmt.Println("DelDataOnIndex error, index < 0, error")
		return
	}

	// 删除头节点
	if index == 0 {
		l.head = l.head.next
		l.length--
		fmt.Println("DelDataOnIndex success, is head")
		return
	}

	i := 1
	cur := l.head
	for cur.next.next != nil {
		if i == index {
			cur.next = cur.next.next
			l.length--
			fmt.Println("DelDataOnIndex success, index:", index)
			return
		}
		i++
		cur = cur.next
	}
	// 到了这里，说明 cur.next.next == nil, 可能要删除最后一个, cur.next是最后一个元素
	if i == index {
		cur.next = cur.next.next
		l.length--
		fmt.Println("DelDataOnIndex success, index:", index)
	} else {
		fmt.Println("DelDataOnIndex error, index not found")
	}
}

func (l *List) DelData(value int) {
	// 先看是不是要删除头节点
	cur := l.head
	if cur.data == value {
		l.head = l.head.next
		l.length--
		fmt.Println("DelData success, value:", value)
		return
	}
	// 找下一个
	for cur.next.next != nil {
		if cur.next.data == value {
			cur.next = cur.next.next
			l.length--
			fmt.Println("DelData success, value:", value)
			return
		}
		cur = cur.next
	}
	// 到了这里，说明 cur.next.next == nil, 可能要删除最后一个, cur.next是最后一个元素
	if cur.next.data == value {
		cur.next = cur.next.next
		l.length--
		fmt.Println("DelData success, value:", value)
	} else {
		fmt.Println("DelData fail, value not found")
	}

}

func (l *List) PrintLink() {
	if l.head == nil {
		fmt.Println("list is nil")
		return
	}
	cur := l.head
	count := 0
	for cur != nil {
		fmt.Println("node data:", cur.data, ", index:", count)
		cur = cur.next
		count++
	}
	fmt.Println("link length:", l.length)
}

func (l *List) FindValue(value int) bool {
	cur := l.head
	for cur != nil {
		if cur.data == value {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *List) UpdateValueByIndex(index int, value int) {
	if index < 0 {
		fmt.Println("UpdateValueByIndex fail, index < 0")
		return
	}
	i := 0
	cur := l.head
	for cur != nil {
		if i == index {
			old := cur.data
			cur.data = value
			fmt.Println("UpdateValueByIndex success,", old, "->", value)
			return
		}
		i++
		cur = cur.next
	}
	fmt.Println("UpdateValueByIndex fail, index not found")
}
