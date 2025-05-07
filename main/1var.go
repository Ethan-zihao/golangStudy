package main

import "fmt"

func main() {
	// 变量声明
	// go中整形和浮点型的默认值都是0，字符串的默认值是空字符串，布尔型的默认值是false，切片的默认值是nil，

	var (
		name string
	)

	name = "John"

	var age int = 30
	// 打印变量
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)
	//匿名变量
	// 匿名变量是一个下划线_，它可以用来忽略一个值，或者忽略一个返回值中的某个值。
	// 匿名变量只能在函数内部使用，不能在函数外部使用。
	// 匿名变量不能被取地址，因为它没有名字。
	// 匿名变量可以被用作占位符，或者忽略一个返回值中的某个值。
	name, _ = getName()
	fmt.Printf("Name: %v\n", name)

}
func getName() (string, int) {
	return "John2", 32
}
