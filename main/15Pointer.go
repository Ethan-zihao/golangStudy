package main
 import (
	"fmt"

)
func main() {
	// 指针是一个变量，其值为另一个变量的地址，即，内存位置的直接地址。
	// 指针变量声明的一般形式为：var var_name *var-type。
	// 指针变量的值是另一个变量的地址。
	// go语言的指针操作非常简单，只需要记住两个符号：&和*。
	// &是取地址符，*是取值符。
	// 指针变量的声明和初始化：
     var a int = 100

	 var p *int
	 fmt.Println(p)  // 0  指针变量的默认值是nil
	//  fmt.Println(*p)  
	 p = &a  // 取地址符
	 fmt.Println(&a)  // 100
	 fmt.Println(p)  // 0xc0000100a8  地址值
	 fmt.Println(*p)  // 100  取值符  指针变量的值是另一个变量的地址。
	//  数组指针：
	// var ptr [MAX] *int;  // 数组指针
	// 
	ArrPointer()

}
func ArrPointer() { 
 a:= []int{100,200,300}  // 数组指针
 fmt.Println(len(a)) 
 // 数组长度
 MAX := len(a)
//  var ptr [3]*int  // 数组指针 []里面的参数必须是常量
 ptr := make([]*int, len(a))
  for i := 0; i < MAX; i++ {
	ptr[i]=&a[i]  // 取地址符
	fmt.Printf("a[%d]=%d\n",i,*ptr[i])  // 取值符  指针变量的值是另一个变量的地址。
  }



}