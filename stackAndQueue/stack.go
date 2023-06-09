package stackAndQueue

import (
	"errors"
	"fmt"
)

/*
栈（Stack）是一种先进后出（Last-In-First-Out，LIFO）的数据结构。它具有两个主要操作：压入（Push）和弹出（Pop）。
在栈中，所有元素都按照一个固定的顺序存储，新元素被添加到栈的顶部，而任何时候只能访问或删除栈顶元素。栈可以用于许多计算机程序中，例如函数调用、表达式求值等等。
*/

type Stack struct {
	elements []interface{}
	size     int
}

type IStack interface {
	StackSize() int             // 获取栈中元素的数量
	IsEmpty() bool              // 栈是否为空
	Push(value interface{})     // 入栈
	Pop() (interface{}, error)  // 弹出栈顶元素并返回其值
	Peek() (interface{}, error) // 返回栈顶元素的值，但不弹出该元素
	Print()                     // 从栈顶开始打印栈中的数据
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

func (s *Stack) StackSize() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
	s.size++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	data := s.elements[s.size-1]
	s.elements = s.elements[:s.size-1]
	s.size--
	return data, nil
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
	}
	for i := s.size - 1; i >= 0; i-- {
		fmt.Println("stack element:", s.elements[i])
	}
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.elements[s.size-1], nil
}
