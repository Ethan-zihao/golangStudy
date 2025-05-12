package main

import "fmt"

// 定于结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	// 格式化输出
	// %v		输出变量的值
	// %#v		输出变量的值和类型
	// %T		输出变量的类型
	// %%		输出%
	zihao := Person{Name: "John", Age: 30}
	fmt.Printf("zihao: %v\n", zihao)
	// 	zihao: {John 30}

	fmt.Printf("zihao.Name: %#v\n", zihao)
	// zihao.Name: main.Person{Name:"John", Age:30}
	fmt.Printf("zihao.Age: %T\n", zihao.Age)
	// zihao.Age: int
	// 布尔占位符
	// %t		输出布尔值
	fmt.Printf("%t\n", true)
	// true
	// 整数占位符
	// %b		输出二进制
	// %d		输出十进制
	// %o		输出八进制
	// %x		输出十六进制
	// %X		输出大写的十六进制
	// %U		输出Unicode字符
	// %c		输出字符
	fmt.Printf("%b\n", 10)
}
