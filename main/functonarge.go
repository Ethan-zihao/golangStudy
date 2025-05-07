package main

import "fmt"

//形参
func sum(a int, b int) int {
	return a + b
}
func f1(a int) {
	a = 100
}
func main() {
	x := 200
	f1(x)
	fmt.Printf("x: %v\n", x)
}
 
