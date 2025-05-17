package main

import "fmt"

func main() {
	// 标记

	test2()
}

// break可以解释for select switch 语句
// 只能跳出一层循环
// 单独在switch中使用break，没有意义
// break后面可以跟标签，跳出到标签处
// 标签名区分大小写，必须定义在对应的for、select、switch的外部
// 在switch中没有fallthrough,使用break没有意义,有flallthrough，使用break可以跳出switch
func test3() {
MYBEARK:
	for i := 0; i < 10; i++ {
		fmt.Println("外部i", i)
		for j := 0; j < 10; j++ {
			if i == 2 {
				//break MYBEARK
				// 跳出当前循环，执行下一个循环
				// 跳出到标记处
				fmt.Println("break", j)
				break MYBEARK
			}
			fmt.Println("continue", j)
		}
	}
}

func test() {
	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		break
		fallthrough // 穿透，执行下一个case
	case 3:
		fmt.Println("3")
	}
}
func test2() {
	// goto 语句
	// 可以跳转到指定的标签处
	// 标签名区分大小写，必须定义在对应的for、select、switch的外部
	i := 2
	if i == 2 {
		goto END
	}
	fmt.Println("i", i)
END:
	fmt.Println("end")

}
