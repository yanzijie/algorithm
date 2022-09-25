package link

import "fmt"

// 双向循环链表
// 这个代码的index是从1开始算

type DoubleNode struct {
	data int
	next *DoubleNode
	pre  *DoubleNode
}

type DoubleListInterface interface {
	InsertHead(value int)                     // 头部插入
	InsertTail(value int)                     // 尾部插入
	InsertByIndex(value int, index int)       // 根据位置插入 index是从1开始算
	DeleteByIndex(index int)                  // 根据位置删除 index是从1开始算
	DeleteByValue(value int)                  // 根据值删除
	UpdateByIndex(index int, newValue int)    // 根据位置修改 index是从1开始算
	UpdateByValue(oldValue int, newValue int) // 根据值删除
	GetIndexByValue(value int) (int, bool)    // 根据值获取位置 只有一种错误可能，就返回bool
	GetValueByIndex(index int) (int, bool)    // 根据位置获取值
	PrintList()
}

type DoubleList struct {
	head   *DoubleNode
	tail   *DoubleNode
	length int
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (d *DoubleList) InsertHead(value int) {
	newData := &DoubleNode{
		data: value,
		next: nil,
		pre:  nil,
	}
	if d.head == nil {
		// 链表没有数据
		d.head = newData
		d.tail = newData
		newData.next = newData
		newData.pre = newData
	} else {
		// 链表有数据
		newData.next = d.head
		newData.pre = d.tail
		d.tail.next = newData
		d.head.pre = newData
		d.head = newData
	}

	d.length++
}

func (d *DoubleList) InsertTail(value int) {
	newNode := &DoubleNode{
		data: value,
		next: nil,
		pre:  nil,
	}
	if d.head == nil {
		// 链表没有数据
		d.head = newNode
		d.tail = newNode
		newNode.next = newNode
		newNode.pre = newNode
	} else {
		// 链表有数据
		newNode.pre = d.tail
		newNode.next = d.head
		d.tail.next = newNode
		d.head.pre = newNode
		d.tail = newNode
	}

	d.length++
}

func (d *DoubleList) InsertByIndex(value int, index int) {
	if index > d.length || index <= 0 {
		fmt.Println("index error")
		return
	}
	newNode := &DoubleNode{
		data: value,
		next: nil,
		pre:  nil,
	}
	if d.head == nil {
		d.InsertHead(value)
		return
	}
	if index == d.length {
		d.InsertTail(value)
	}

	cur := d.head
	// cur指向头节点的时候, count=1
	count := 1
	for cur != d.tail {
		if count == index {
			newNode.next = cur
			newNode.pre = cur.pre
			cur.pre.next = newNode
			cur.pre = newNode
			d.length++
		}
		cur = cur.next
		count++
	}

}

func (d *DoubleList) DeleteByIndex(index int) {
	if d.head == nil {
		fmt.Println("link nil")
		return
	}
	if index > d.length || index <= 0 {
		fmt.Println("index error")
		return
	}

	if d.head == d.tail {
		// 只有一个节点
		d.head = nil
		d.tail = nil
		d.length = 0
		return
	}

	cur := d.head
	// cur指向头节点的时候, count=1
	count := 1
	fmt.Println("DeleteByIndex:", index)
	// 当cur转了一圈回到头节点时候，count肯定不等于1，退出循环
	for cur != d.head || count == 1 {
		if count == index {
			cur.pre.next = cur.next
			cur.next.pre = cur.pre
			if index == d.length {
				// 重新指定尾节点
				d.tail = cur.pre
			}
			if index == 1 {
				// 重新指定头节点
				d.head = cur.next
			}
			// 删除当前节点的前后指向关系，让它与链表失去关联，则会被GC回收
			cur.next = nil
			cur.pre = nil
			d.length--
			return
		}
		cur = cur.next
		count++
		if count > d.length {
			// 说明已经遍历完链表了
			break
		}
	}
	fmt.Println("DeleteByIndex error")
}

func (d *DoubleList) DeleteByValue(value int) {
	if d.head == nil {
		fmt.Println("link is nil")
		return
	}

	cur := d.head
	count := 1
	fmt.Println("DeleteByValue:", value)
	// 当cur转了一圈回到头节点时候，count肯定不等于1，退出循环
	for cur != d.head || count == 1 {
		if value == cur.data {
			cur.pre.next = cur.next
			cur.next.pre = cur.pre
			if count == d.length {
				// 重新指定尾节点
				d.tail = cur.pre
			}
			if count == 1 {
				// 重新指定头节点
				d.head = cur.next
			}
			// 删除当前节点的前后指向关系，让它与链表失去关联，则会被GC回收
			cur.next = nil
			cur.pre = nil
			d.length--
			return
		}
		cur = cur.next
		count++
		if count > d.length {
			// 说明已经遍历完链表了
			break
		}
	}
	fmt.Println("DeleteByValue error")
}

func (d *DoubleList) UpdateByIndex(index int, newValue int) {
	if index > d.length || index <= 0 {
		fmt.Println("index error")
		return
	}

	cur := d.head
	count := 1
	fmt.Println("UpdateByIndex:", index, ", change to:", newValue)
	// 当cur转了一圈回到头节点时候，count肯定不等于1，退出循环
	for cur != d.head || count == 1 {
		if count == index {
			cur.data = newValue
			return
		}
		cur = cur.next
		count++
		if count > d.length {
			// 说明已经遍历完链表了
			break
		}
	}
	fmt.Println("UpdateByIndex error")
}

func (d *DoubleList) UpdateByValue(oldValue int, newValue int) {
	cur := d.head
	count := 1
	fmt.Println("UpdateByValue oldValue:", oldValue, ", change to:", newValue)
	// 当cur转了一圈回到头节点时候，count肯定不等于1，退出循环
	for cur != d.head || count == 1 {
		if oldValue == cur.data {
			cur.data = newValue
			return
		}
		cur = cur.next
		count++
		if count > d.length {
			// 说明已经遍历完链表了
			break
		}
	}
	fmt.Println("UpdateByValue error")
}

func (d *DoubleList) GetIndexByValue(value int) (int, bool) {
	cur := d.head
	count := 1
	// 当cur转了一圈回到头节点时候，count肯定不等于1，退出循环
	for cur != d.head || count == 1 {
		if value == cur.data {
			// 找到了
			return count, true
		}
		cur = cur.next
		count++
		if count > d.length {
			// 说明已经遍历完链表了
			break
		}
	}
	return 0, false
}

func (d *DoubleList) GetValueByIndex(index int) (int, bool) {
	if index > d.length || index <= 0 {
		return 0, false
	}

	cur := d.head
	count := 1
	// 当cur转了一圈回到头节点时候，count肯定不等于1，退出循环
	for cur != d.head || count == 1 {
		if index == count {
			return cur.data, true
		}
		cur = cur.next
		count++
		if count > d.length {
			// 说明已经遍历完链表了
			break
		}
	}
	return 0, false
}

func (d *DoubleList) PrintList() {
	if d.head == nil {
		fmt.Println("list is null")
		return
	}

	cur := d.head
	count := 1
	for {
		//先打印再判断
		fmt.Println("node data:", cur.data, ",index:", count,
			",next data:", cur.next.data,
			",pre data:", cur.pre.data)
		if cur.next == d.head {
			break
		}
		cur = cur.next
		count++
	}
	fmt.Println("list length =", d.length)
}
