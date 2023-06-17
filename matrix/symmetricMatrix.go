package matrix

// 对称矩阵的压缩存储, 矩阵下标从1开始算
/* a(i,j)
a(1,1)  a(1,2)  a(1,3)  a(1,4)
a(2,1)  a(2,2)  a(2,3)  a(2,4)
a(3,1)  a(3,2)  a(3,3)  a(3,4)
a(4,1)  a(4,2)  a(4,3)  a(4,4)
*/

type SymmetricMatrix struct {
	size int   // 矩阵的阶数 (就是n)
	data []int // 存储压缩后的上三角元素或下三角元素
}

func NewSymmetricMatrix(size int) *SymmetricMatrix {
	return &SymmetricMatrix{
		size: size,
		data: make([]int, (size*(size+1))/2), // n(n+1)/2	压缩后的一维数组长度
	}
}

type ISymmetricMatrix interface {
	Get(i, j int) int    // 获取矩阵中某个位置上的元素
	Set(i, j, value int) // 设置矩阵中某个位置上的元素
	Compress()           // 压缩存储矩阵
	Decompress()         // 解压缩矩阵
}

func (s *SymmetricMatrix) Get(i, j int) int {
	// 矩阵下标从1开始算
	if i > s.size || j > s.size || i < 1 || j < 1 {
		panic("Index out of range")
	}
	if i < j {
		i, j = j, i // 转换为上三角矩阵
	}
	// i*(i-1)/2 + j - 1  是这个矩阵数据在压缩后的一维数组中的下标
	k := i*(i-1)/2 + j - 1
	return s.data[k]
}

func (s *SymmetricMatrix) Set(i, j, value int) {
	// 矩阵下标从1开始算
	if i > s.size || j > s.size || i < 1 || j < 1 {
		panic("Index out of range")
	}
	if i < j {
		i, j = j, i // 转换为上三角矩阵
	}
	// i*(i-1)/2 + j - 1  是这个矩阵数据在压缩后的一维数组中的下标
	k := i*(i-1)/2 + j - 1
	s.data[k] = value
}

func (s *SymmetricMatrix) Compress() {
	// 遍历上三角部分，将非零元素保存在数组中
	index := 0
	for i := 1; i <= s.size; i++ {
		for j := i; j <= s.size; j++ {
			s.data[index] = s.Get(i, j)
			index++
		}
	}
}

func (s *SymmetricMatrix) Decompress() {
	// 将数组中的元素还原到矩阵中
	index := 0
	for i := 1; i <= s.size; i++ {
		for j := i; j <= s.size; j++ {
			s.Set(i, j, s.data[index])
			index++
		}
	}
}
