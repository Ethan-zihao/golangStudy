package main

import "fmt"

func main() {
	// f1()
	f3()
}

// func f1() {
// 	count := 10
// 	for i := 0; i < count; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}
// }

// 初始条件卸在外面
// func f2() {
// 	count := 10
// 	i := 0
// 	for ; i < count; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}

// }

// 初始条件和结束条件都省略

func f3() {
	i := 1
	for i < 10 {
		fmt.Printf("i: %v\n", i)
		   i++
	}

}
