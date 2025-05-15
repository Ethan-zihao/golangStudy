package main

	//   闭包
	//   闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
	//   闭包的价值在于可以作为函数对象或者匿名函数，对于类型系统而言，这意味着不仅要表示数据还要表示代码。
	// code
	import "fmt"
	func add() func(int) int{
		sum:=0
		return func(x int) int{
			fmt.Printf("sumstart: %v\n", sum)
			sum+=x
			return sum
		}
	}
func main() {
      f:=add()
	  r:=f(10)
	  fmt.Println(r)
	  r=f(20)
	  fmt.Println(r)
	  r=f(30)
	  fmt.Println(r)
}
