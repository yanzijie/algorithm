package link

import (
	"fmt"
	"sync"
)

// 单向链表

type IOneWayLinkList interface {
	InsertOnHead(value int)             //从链表头部插入
	InsertOnEnd(value int)              //从尾部添加元素
	InsertOnIndex(value int, index int) //在指定位置添加元素  index是下标，从0开始
	DelDataOnIndex(index int)           //删除链表指定位置的节点, index是下标，从0开始
	DelData(value int)                  //删除链表指定值的节点
	PrintLink()                         //打印链表所有节点
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
	lock   sync.Mutex
}

func InitOneWayLinkList() IOneWayLinkList {
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
		fmt.Println("InsertOnHead success, this head")
		return
	}
	n.next = l.head
	l.head = n
	fmt.Println("InsertOnHead success")
}

func (l *List) InsertOnEnd(value int) {
	//TODO implement me
	panic("implement me")
}

func (l *List) InsertOnIndex(value int, index int) {
	//TODO implement me
	panic("implement me")
}

func (l *List) DelDataOnIndex(index int) {
	//TODO implement me
	panic("implement me")
}

func (l *List) DelData(value int) {
	//TODO implement me
	panic("implement me")
}

func (l *List) PrintLink() {
	if l.head == nil {
		fmt.Println("list is nil")
		return
	}
	cur := l.head
	fmt.Printf("list : ")
	for cur != nil {
		fmt.Printf("%d ", cur.data)
		cur = cur.next
	}
	fmt.Println()
}
