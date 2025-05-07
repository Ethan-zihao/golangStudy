package main

import (
	"fmt"
	"runtime"
)

func TypeTest() *int {
	a := 100
	p := &a
	return p

}
func main() {
	fmt.Println("wadwad")
	var a1 = [3]int{12, 13}
	a1[2] = 14
	a := 100
	age := 19
	if age >= 18 {
		fmt.Println("你还未成年")
	} else {
		fmt.Println("你还未成年")
	}
	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Printf("i:%v\n", i)
	// }

	p := &a
	b3 := false
	WA := [...]int{1, 3, 4}
	// for i, v := range WA {
	// 	fmt.Printf("WA[%v]: %v\n", i, v)

	// }
	// for _, v := range WA {
	// 	fmt.Printf("V:%v\n ", v)
	// }
	for _, v := range WA {
		fmt.Printf("v: %v\n", v)
	}
	// as:=[2]int{1, 2,}
	as := []int{1, 23}
	fmt.Printf("%T\n", WA)
	fmt.Printf("%T\n", as)
	fmt.Printf("%T\n", p)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("2: %T\n", b3)
	fmt.Printf("runtime.Version(): %v\n", runtime.Version())

}
