package main

import (
	"fmt"
)

func main() {

	test()
}
func test() {
	// 声明切片
	var s2 = make([]int, 2)
	var s1 []int
	fmt.Printf("s2: %v\n", s2)

	fmt.Printf("s1: %v\n", s1)
	var s3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("cap(s1): %v\n", cap(s3))
	fmt.Printf("len(s3): %v\n", len(s3))
	fmt.Printf("s1[1]: %v\n", s3[1])
	//切片初始化
	// 数组截取【0-3】作闭右开    [:]表示全部
	// 一个切片未初始化前默认是nil长队为0，容量为0
	// 空切片 【nil】 长度为0，容量为0

	s4 := s3[0:3]
	fmt.Printf("s4: %v\n", s4)
	s5 := s3[3:]
	fmt.Printf("s5: %v\n", s5)
	s6 := s3[:]
	fmt.Printf("s6: %v\n", s6)
	var s7 []int

	fmt.Printf("(s7 == nil): %v\n", (s7 == nil))
	s8 := s3[:3]
	fmt.Printf("s8: %v\n", s8)
}
