package testType

import "fmt"

func TypeTest() *int {
	a := 100
	// 创建一个指针变量 p，指向变量 a
	p := &a
	fmt.Printf("p: %v\n", p)
	// 定义数组
	arr := [3]int{1, 2, 3}
	fmt.Printf("arr: %v\n", arr)
 return &a
}
func main() {
	p:=TypeTest()
	fmt.Printf("p: %v\n", p)
}