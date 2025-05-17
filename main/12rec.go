package main
 import "fmt"
func main() {
	// 递归函数
	// 递归函数是指自己调用自己的函数。递归函数在解决许多数学问题上都有广泛的应用，因此在学习Go语言的过程中，我们也应该掌握递归函数的使用。
	// 递归函数的使用需要注意以下几点：
	// 1. 递归函数必须有一个终止条件，否则会导致无限递归，最终导致栈溢出。
	// 2. 递归函数必须有一个递推关系，即如何从当前状态转换到下一个状态	。
	// 3. 递归函数必须有一个返回值，即如何从下一个状态返回上一个状态。
	// 4. 递归函数的效率比较低，因为每次调用函数都会占用栈空间，因此在使用递归函数时，需要注意栈空间的使用。
	// 实例 ：
	// 求阶乘
	// 阶乘是一个数学概念，它表示从1到n的所有整数的乘积。例如，5的阶乘是1*2*3*4*5=120
	factorial := factorial(5)
	fmt.Println(factorial)
	// 斐波那契数列
	// 斐波那契数列是一个数学概念，它表示从1到n的所有整数的和。例如，5的斐波那契数列是1+2+3+4+5=15
    fibonaci := fibonaci(5)
	fmt.Println(fibonaci)
}
// code
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
func fibonaci(n int)int{
	if n==0{
		return 0
	}else if n==1{
		return 1 
	}else{
		return fibonaci(n-1)+fibonaci(n-2)
	}
}