package matrix

import "fmt"

// Node 稀疏矩阵节点结构体
type Node struct {
	row   int // 节点所在行数
	col   int // 节点所在列数
	value int // 节点的值
}

// SparseMatrix 稀疏矩阵结构体
type SparseMatrix struct {
	rows  int    // 矩阵的行数
	cols  int    // 矩阵的列数
	nodes []Node // 矩阵的非零节点数组
}

// ConvertToSparseMatrix 将普通矩阵转换为稀疏矩阵
func ConvertToSparseMatrix(matrix [][]int) SparseMatrix {
	var sparseMatrix SparseMatrix
	sparseMatrix.rows = len(matrix)
	sparseMatrix.cols = len(matrix[0])
	for i := 0; i < sparseMatrix.rows; i++ {
		for j := 0; j < sparseMatrix.cols; j++ {
			if matrix[i][j] != 0 {
				var node Node
				node.row = i
				node.col = j
				node.value = matrix[i][j]
				sparseMatrix.nodes = append(sparseMatrix.nodes, node)
			}
		}
	}
	return sparseMatrix
}

// ConvertToNormalMatrix 将稀疏矩阵转换为普通矩阵
func ConvertToNormalMatrix(sparseMatrix SparseMatrix) [][]int {
	var normalMatrix [][]int
	for i := 0; i < sparseMatrix.rows; i++ {
		row := make([]int, sparseMatrix.cols)
		normalMatrix = append(normalMatrix, row)
	}
	for _, node := range sparseMatrix.nodes {
		// 每个节点包含着所在的行和列, 所以这里直接使用行和列再赋值就行了
		normalMatrix[node.row][node.col] = node.value
	}
	return normalMatrix
}

// Print 输出稀疏矩阵
func (sparseMatrix SparseMatrix) Print() {
	fmt.Println("压缩后的稀疏矩阵:")
	fmt.Printf("Rows: %d, Cols: %d\n", sparseMatrix.rows, sparseMatrix.cols)
	for _, node := range sparseMatrix.nodes {
		fmt.Printf("(%d, %d)=%d ", node.row, node.col, node.value)
	}
	fmt.Println()
}
