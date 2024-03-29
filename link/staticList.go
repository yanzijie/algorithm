package link

// 静态链表
/*
1.定义: 使用数组存储的链表,就是分配一片空间，链表中的元素都存放在这个空间中,
链表中的元素有两个成员:
	数据域：存储节点的数据信息
	指针域：下一个节点在数组中的位置下标
2.优点: 不需要像动态链表那样频繁地进行内存分配和释放，因此可以减少内存碎片的产生，提高空间利用率。此外，在某些场景下，静态链表也可以避免由于频繁的内存分配和释放导致的性能问题
3.缺点: 插入和删除节点时需要移动大量元素，因此操作效率较低。

代码略.....
*/
