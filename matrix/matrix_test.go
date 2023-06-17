package matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSymmetricMatrix(t *testing.T) {
	// 对称矩阵压缩存储测试
	/*
		a(1,1)  a(1,2)  a(1,3)  a(1,4)
		a(2,1)  a(2,2)  a(2,3)  a(2,4)
		a(3,1)  a(3,2)  a(3,3)  a(3,4)
		a(4,1)  a(4,2)  a(4,3)  a(4,4)
	*/

	// 初始化一个矩阵
	m := NewSymmetricMatrix(3)
	m.Set(1, 1, 1)
	m.Set(2, 1, 2)
	m.Set(2, 2, 3)
	m.Set(3, 1, 4)
	m.Set(3, 2, 5)
	m.Set(3, 3, 6)

	// 测试 Get 方法
	if m.Get(1, 1) != 1 || m.Get(2, 1) != 2 || m.Get(2, 2) != 3 ||
		m.Get(3, 1) != 4 || m.Get(3, 2) != 5 || m.Get(3, 3) != 6 {
		t.Errorf("Get method failed")
	}

	// 压缩矩阵并测试结果是否正确
	m.Compress()
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(m.data, expected) {
		t.Errorf("Compress method failed")
	}

	// 解压缩矩阵并测试结果是否正确
	m.Decompress()
	for i := 1; i <= m.size; i++ {
		for j := 1; j <= m.size; j++ {
			if m.Get(i, j) != i+j-1 {
				t.Errorf("Decompress method failed")
			}
		}
	}
}

// 测试稀疏矩阵压缩存储
func TestSparseMatrix(t *testing.T) {
	// 测试数据
	normalMatrix := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 2, 0},
		{0, 0, 0, 0, 0},
		{0, 3, 0, 4, 0},
		{0, 0, 0, 0, 0},
	}

	// 将普通矩阵转换为稀疏矩阵
	sparseMatrix := ConvertToSparseMatrix(normalMatrix)
	// 输出稀疏矩阵
	fmt.Println("稀疏矩阵行数:", sparseMatrix.rows)
	fmt.Println("稀疏矩阵列数:", sparseMatrix.cols)
	fmt.Println("稀疏矩阵内的值:", sparseMatrix.nodes)

	// 输出稀疏矩阵
	sparseMatrix.Print()

	// 将稀疏矩阵转换为普通矩阵
	normalMatrix = ConvertToNormalMatrix(sparseMatrix)

	fmt.Println("还原为普通矩阵")
	// 输出普通矩阵
	for _, row := range normalMatrix {
		fmt.Println(row)
	}
}
