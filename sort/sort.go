package sort

import "fmt"

// BubbleSort 冒泡排序 [16,4,12,5,9,14,56,10,35,23], 升序排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		isSwap := false // 是否发生交换
		// 从后面往前，把小的数往最前面放
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
				isSwap = true
			}
		}
		if !isSwap {
			// 没有发生交换，说明已经有序了
			break
		}
	}
	fmt.Println("after sort:", arr)
}

// SelectSort 选择排序 [16,4,12,5,9,14,56,10,35,23]
func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i // 最小元素的下标
		// 找到最小的位置
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		// 找到最小的下标位置了，交换
		if min != i {
			tmp := arr[min]
			arr[min] = arr[i]
			arr[i] = tmp
		}
	}
	fmt.Println("after SelectSort arr:", arr)
}

// QuickSort 快速排序 [16,4,12,5,9,14,56,10,35,23]
func QuickSort(arr []int, low int, high int) []int {
	if low < high {
		// 划分数组
		p := partition(arr, low, high)
		// 划分左子表
		QuickSort(arr, low, p-1)
		// 划分右子表
		QuickSort(arr, p+1, high)
	}
	return arr
}

//划分数组 [16,4,12,5,9,14,56,10,35,23]
func partition(arr []int, low int, high int) int {
	p := arr[low]
	for low < high {
		for low < high && arr[high] >= p {
			// >=基准元素, 游标左移
			high--
		}
		// 到这里说明arr[high] < p, 移动元素
		arr[low] = arr[high]
		for low < high && arr[low] < p {
			// <基准元素，游标移动
			low++
		}
		// 到这里说明arr[high] >= p, 移动元素
		arr[high] = arr[low]
	}
	// 到这说明 low == high 移动基准元素到这个位置
	arr[low] = p
	return low
}
