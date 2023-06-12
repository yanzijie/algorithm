package stackAndQueue

import (
	"errors"
	"fmt"
)

/*
队列是一个有序列表，可以用数组或是链表来实现,
遵循先入先出的原则。即：先存入队列的数据，要先取出。后存入的要后取出

这里使用切片实现一个双端队列
[1,2,3,4,5]   1是队头, 5是队尾
*/

// Deque 双端队列
type Deque struct {
	items []int // 队列中的数据
	size  int   // 队列长度, [1,2,3,4] 这种长度算4
}

type IDeque interface {
	AddFront(item int)         // 在队列头部添加
	RemoveRear() (int, error)  // 在队列尾部返回然后删除
	AddRear(item int)          // 在队列尾部添加
	RemoveFront() (int, error) // 在队列头部返回然后删除
	IsEmpty() bool             // 队列是否为空
	Print()                    // 从队列头部开始打印
}

func NewDeque() *Deque {
	return &Deque{
		items: make([]int, 0),
		size:  0,
	}
}

func (d *Deque) AddFront(item int) {
	// 搞一个新切片[]int{item}内含要添加的元素item, 然后把原来的切片添加进去,就在新元素后面了，再赋值给原来的切片, 6
	d.items = append([]int{item}, d.items...)
	d.size++
}

func (d *Deque) RemoveRear() (int, error) {
	if d.IsEmpty() {
		fmt.Println("queue is empty")
		return 0, errors.New("queue is empty")
	}
	item := d.items[d.size-1]
	d.items = d.items[:d.size-1]
	d.size--
	return item, nil
}

func (d *Deque) AddRear(item int) {
	d.items = append(d.items, item)
	d.size++
}

func (d *Deque) RemoveFront() (int, error) {
	if d.IsEmpty() {
		fmt.Println("queue is empty")
		return 0, errors.New("queue is empty")
	}
	item := d.items[0]
	d.items = d.items[1:]
	d.size--
	return item, nil
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) Print() {
	if d.IsEmpty() {
		fmt.Println("queue is empty")
		return
	}
	for i := 0; i < d.size; i++ {
		fmt.Println("queue item:", d.items[i])
	}
}
