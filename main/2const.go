package main

import "fmt"

func main() {
	// iota  不能使用:= 
	// iota 是一个常量计数器，默认从0开始，每遇到一个const关键字，iota的值就会+1
	//  _ 表示忽略这个值
	const (
		a = iota // iota = 0
		b = iota // iota = 1
		_        // iota = 	2
		c = iota // iota = 3
		d        // iota = 4
	)
	const conName = "John"
	const PI float32 = 3.14
	const PI2 = 3.14159
	fmt.Printf("conName: %v\n", conName)
}