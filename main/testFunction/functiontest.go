package main

import "fmt"

func test(a int, b int) (ret int) {
	ret = a + b
	fmt.Printf("ret: %v\n", ret)
	return ret
}
func test3() (string, int) {
	n := "to"
	a := 13
	return n, a
}
func main() {
	test(10, 2)
	name, age := test3()
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

}
