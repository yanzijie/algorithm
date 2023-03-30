package link

import "fmt"

// 单向循环链表

type OneLoopNode struct {
	data int
	next *OneLoopNode
}

type OneLoopLink struct {
	head   *OneLoopNode
	tail   *OneLoopNode
	length int
}

type OneLoopLinkInterface interface {
	InsertToHead(value int)                    // 添加元素到链表头部
	InsertToTail(value int)                    // 添加元素到链表尾部
	InsertByIndex(index, value int) bool       // 添加元素在指定位置  index是下标，从0开始,
	DeleteByIndex(index int) bool              // 删除链表指定位置的节点, index是下标，从0开始
	DeleteByValue(value int) bool              // 删除链表指定值的节点
	ModifyByIndex(index, value int) bool       // 修改 index 下标的值为value, 下标从0开始
	ModifyByValue(oldValue, newValue int) bool // 修改 oldValue值为 newValue
	SearchByIndex(index int) (int, bool)       // 查询 index 下标对应节点value值, index是下标，从0开始
	SearchByValue(value int) (int, bool)       // 查询 value 值对应节点的 index 下标,从0开始
	PrintfLink()                               // 打印链表所有数据
}

func NewOneLoopLink() OneLoopLinkInterface {
	return &OneLoopLink{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (o *OneLoopLink) InsertToHead(value int) {
	newNode := &OneLoopNode{}
	newNode.data = value
	newNode.next = nil

	if o.head == nil {
		o.head = newNode
		o.tail = newNode
		o.tail.next = o.head
		o.length++
		return
	}

	newNode.next = o.head
	o.tail.next = newNode
	o.head = newNode
	o.length++

	fmt.Println("InsertToHead success")
}

func (o *OneLoopLink) InsertToTail(value int) {
	newNode := &OneLoopNode{}
	newNode.data = value
	newNode.next = nil

	if o.head == nil {
		o.head = newNode
		o.tail = newNode
		o.tail.next = o.head
		o.length++
		return
	}

	o.tail.next = newNode
	o.tail = newNode
	newNode.next = o.head
	o.length++

	fmt.Println("InsertToTail success")
}

func (o *OneLoopLink) InsertByIndex(index, value int) bool {
	if index == 0 || o.head == nil {
		o.InsertToHead(value)
		return true
	}
	if index == o.length-1 {
		o.InsertToTail(value)
		return true
	}
	if index >= o.length {
		fmt.Println("InsertByIndex fail")
		return false
	}

	newNode := &OneLoopNode{}
	newNode.data = value
	newNode.next = nil
	cur := o.head.next
	pre := o.head
	count := 1
	// 先干活再判断
	for {
		if count == index {
			newNode.next = cur
			pre.next = newNode
			o.length++
			fmt.Println("InsertByIndex success")
			return true
		}
		if cur.next == o.head {
			break
		}
		cur = cur.next
		pre = pre.next
		count++
	}
	return false
}

func (o *OneLoopLink) DeleteByIndex(index int) bool {
	if index >= o.length {
		fmt.Println("DeleteByIndex fail")
		return false
	}
	if o.head == o.tail {
		// 只有一个节点
		o.head = nil
		o.tail = nil
		o.length = 0
		fmt.Println("DeleteByIndex success")
		return true
	}
	if index == 0 {
		// 删除头节点
		o.tail.next = o.head.next
		o.head = o.head.next
		o.length--
		fmt.Println("Delete head success, ")
		return true
	}
	cur := o.head.next
	pre := o.head
	count := 1
	for {
		if count == index {
			if index == o.length-1 {
				// 删除尾节点
				pre.next = cur.next
				o.tail = pre
				o.length--
				fmt.Println("Delete tail success")
				return true
			}
			pre.next = cur.next
			o.length--
			fmt.Println("DeleteByIndex success,index=", index)
			return true
		}
		if cur == o.tail {
			break
		}
		cur = cur.next
		pre = pre.next
		count++
	}

	fmt.Println("DeleteByIndex fail")
	return false
}

func (o *OneLoopLink) DeleteByValue(value int) bool {
	if o.head == o.tail {
		// 只有一个节点
		o.head = nil
		o.tail = nil
		o.length = 0
		fmt.Println("DeleteByValue success, value=", value)
		return true
	}

	if o.head.data == value {
		o.tail.next = o.head.next
		o.head = o.head.next
		o.length--
		fmt.Println("DeleteByValue success, value=", value)
		return true
	}

	cur := o.head.next
	pre := o.head
	for {
		if cur.data == value {
			pre.next = cur.next
			o.length--
			fmt.Println("DeleteByValue success, value=", value)
			return true
		}
		if cur == o.tail {
			// 绕到尾巴了还没找到,退出
			break
		}
		cur = cur.next
		pre = pre.next
	}

	return false
}

func (o *OneLoopLink) ModifyByIndex(index, value int) bool {
	if index >= o.length {
		fmt.Println("index too long...")
		return false
	}

	// 先把头节点的情况处理了
	if index == 0 {
		fmt.Println("ModifyByIndex success, index:", index, " value:", o.head.data, " change to:", value)
		o.head.data = value
		return true
	}

	cur := o.head.next
	pre := o.head
	num := 1
	for {
		if num == index {
			fmt.Println("ModifyByIndex success, index:", num, " value:", cur.data, " change to:", value)
			cur.data = value
			return true
		}
		if num >= o.length {
			// 没找到
			break
		}
		cur = cur.next
		pre = pre.next
		num++
	}
	return false
}

func (o *OneLoopLink) ModifyByValue(oldValue, newValue int) bool {
	if o.head.data == oldValue {
		fmt.Println("ModifyByValue success, oldValue:", o.head.data, " change to:", newValue)
		o.head.data = newValue
		return true
	}
	cur := o.head.next
	pre := o.head
	for {
		if cur.data == oldValue {
			cur.data = newValue
			fmt.Println("ModifyByValue success, oldValue:", cur.data, " change to:", newValue)
			return true
		}
		if cur == o.tail {
			break
		}
		cur = cur.next
		pre = pre.next
	}
	return false
}

func (o *OneLoopLink) SearchByIndex(index int) (int, bool) {
	if index < 0 || index > o.length {
		fmt.Println("index error")
		return 0, false
	}
	if index == 0 {
		fmt.Println("SearchByIndex success, index:", index, " value is:", o.head.data)
		return o.head.data, true
	}
	if index == o.length-1 {
		fmt.Println("SearchByIndex success, index:", index, " value is:", o.tail.data)
		return o.tail.data, true
	}
	p := 1
	cur := o.head.next
	for cur != o.tail {
		if p == index {
			fmt.Println("SearchByIndex success, index:", index, " value is:", cur.data)
			return cur.data, true
		}
		cur = cur.next
		p++
	}

	fmt.Println("SearchByIndex failed, index:", index)
	return 0, false
}

func (o *OneLoopLink) SearchByValue(value int) (int, bool) {
	if o.head.data == value {
		fmt.Println("SearchByValue success, value:", value, " index is: 0", o.head.data)
		return o.head.data, true
	}
	if o.tail.data == value {
		fmt.Println("SearchByValue success, value:", value, " index is: ", o.length-1)
		return o.tail.data, true
	}
	cur := o.head.next
	index := 1
	for cur != o.tail {
		if cur.data == value {
			fmt.Println("SearchByValue success, value:", value, " index is:", index)
			return cur.data, true
		}
		cur = cur.next
		index++
	}

	fmt.Println("SearchByValue failed, value:", value)
	return 0, false
}

func (o *OneLoopLink) PrintfLink() {
	if o.head == nil {
		fmt.Println("link is nil")
		return
	}

	cur := o.head
	count := 0
	//先打印再判断
	for {
		fmt.Println("index=", count, ", value=", cur.data, ", next=", cur.next.data)
		if cur.next == o.head {
			break
		}
		cur = cur.next
		count++
	}
	fmt.Println("link length=", o.length)
}
