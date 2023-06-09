package stackAndQueue

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	/*
		1.创建一个新的空栈，检查是否为空和大小是否为零。
		2.尝试从空栈中弹出元素，并检查是否返回预期的错误。
		3.尝试从空栈中获取栈顶元素，并检查是否返回预期的错误。
		4.添加三个值到栈中，检查其大小是否正确。
		5.获取栈顶元素并验证它是预期值。
		6.从栈中弹出一个元素并验证它是预期值。
		7.检查栈的大小是否正确。
		8.打印栈中的数据，以确保它们按正确的顺序打印。
	*/

	stack := NewStack()

	if !stack.IsEmpty() {
		t.Error("expected new stack to be empty")
	}

	if stack.StackSize() != 0 {
		t.Errorf("expected stack size to be %d but got %d", 0, stack.StackSize())
	}

	_, err := stack.Pop()
	if err == nil || err.Error() != "stack is empty" {
		t.Errorf("expected Pop() to return error 'stack is empty'")
	}

	_, err = stack.Peek()
	if err == nil || err.Error() != "stack is empty" {
		t.Errorf("expected Peek() to return error 'stack is empty'")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println("push data done")
	stack.Print()

	if stack.StackSize() != 5 {
		t.Errorf("expected stack size to be %d but got %d", 3, stack.StackSize())
	}

	top, err := stack.Peek()
	if err != nil {
		t.Errorf("expected Peek() to not return an error but got %s", err.Error())
	}
	if top != 5 {
		t.Errorf("expected Peek() to return %d but got %v", 3, top)
	}

	val, err := stack.Pop()
	if err != nil {
		t.Errorf("expected Pop() to not return an error but got %s", err.Error())
	}
	if val != 5 {
		t.Errorf("expected Pop() to return %d but got %v", 3, val)
	}

	if stack.StackSize() != 4 {
		t.Errorf("expected stack size to be %d but got %d", 2, stack.StackSize())
	}

	fmt.Println("after option")
	stack.Print()
}
