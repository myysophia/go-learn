package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

type slice struct {
	array unsafe.Pointer // 用来存储实际数据的数组指针，指向一块连续的内存
	len   int            // 切片中元素的数量
	cap   int            // array数组的长度
}

func main() {
	//go的slice在函数参数传递的时候是值传递还是引用传递： 值传递， 效果又呈现出了引用的效果（不完全是）
	//var data []int
	//for i:=0;i<2000;i++ {
	//	data = append(data, i)
	//	fmt.Printf("len: %d, cap: %d\r\n", len(data), cap(data)) // cap 扩容的规则是：如果当前的cap小于1024，那么每次扩容都是上一次的2倍，如果大于1024，那么每次扩容都是上一次的1.25倍
	//}
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := data[1:6]
	s2 := data[2:7]

	fmt.Println(len(s1), cap(s1)) //  len: 5, cap: 9 默认cap是从切片的起始位置到底层数组的末尾
	fmt.Println(len(s2), cap(s2)) // 	len: 5, cap: 8 从2开始到底层数组的末尾
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
	}
	s1[1] = 22 // 修改s1的值，s2的值也会跟着改变，因为他们共享了底层数组，但是当s2的cap不够的时候，会重新分配一个新的底层数组，这个时候s1和s2就不再共享底层数组了. 修改底层数组就不会影响s2的值了
	fmt.Println(s1)
	fmt.Println(s2)

}
