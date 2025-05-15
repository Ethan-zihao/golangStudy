package main
import "fmt"

func main() {
	// go函数包括：普通函数、匿名函数、闭包函数
	// go杭州不允许嵌套函数，但是允许匿名函数和闭包函数
	// go不允许函数重载，但是允许函数可变参数
    // go中经常使用其中一个返回值作为函数是否执行成功，例如return value.exists,return value.ok,return value.err
    //    参数
    //    go语言是通过传值来传递参数的，即函数参数是值的拷贝。
    //    go语言中函数的参数传递方式有两种：值传递和引用传递。
    //    值传递：函数参数是值的拷贝，函数内部对参数的修改不会影响到函数外部的变量。
    //    引用传递：函数参数是变量的地址，函数内部对参数的修改会影响到函数外部的变量。
	//   像slice、map、channel、interface等类型都是引用类型，它们的默认传递方式是引用传递。
     var arr =[5]int{1,2,3,4,5}
     var slice=[]int{1,2,3,4,5}
	 var m=map[string]int{"apple":1,"banana":2,"orange":3}
	//  test(arr)
	// f3(arr)  //报错f3 函数声明为接收可变参数 ...int，这意味着它期望接收多个单独的 int 类型参数，而不是一个数组或切片
    // arr 是一个 [5]int 类型的数组，不是多个 int 参数。要将数组作为参数传递给 f3 函数，需要将数组变为切片再...展开

    f3(arr[:]...) //
	// f3(1,2,3,4,5)
	test(slice)
	test(arr)
	test(m)
	// 在 Go 中，字符会被自动转换为对应的 ASCII 码值：
	// 字符 '3' 的 ASCII 码值是 51。
	// count:=f4('3',2)   	//  count: 53
	count:=f4(3,2)   	//  count: 5
	fmt.Printf("count: %v\n", count)
	// 函数作为参数
	test2("zhangsan",sayHello)
	ff := calculator("+")(1,2)
	fmt.Printf("ff: %v\n", ff)
	// 匿名函数
    // 匿名函数是指没有函数名的函数，它可以作为参数传递给其他函数，或者直接赋值给变量。
     // 匿名函数可以没有返回值
	func(a int, b int) int {
	    max := 0
	    if a > b {
	        max = a	
	    }else{
	        max = b
		}
		   fmt.Printf("max: %v\n", max)
	}(1, 2)
	// 自己执行

}
// // 统一遍历函数 ，参数可能是数组、切片、map、结构体、接口F
func test (data interface{}){
	fmt.Printf("data: %v\n", data)
	switch T:= data.(type){
	    case[]int:
		   for i,v :=range T{
	         fmt.Printf("slice[%d] = %d\n", i, v)
		   }
		case map[string]int:
		   for key,value:=range T{
			fmt.Printf("map[%s] = %d\n", key, value)
	    }
        case [5]int:
		   for i,v :=range T{
	        fmt.Printf("arr[%d] = %d\n", i, v)	
	   }
	
	}	
}
// 变长参数
func f3(args ...int){
	for _, v:=range args{
		fmt.Println("%d",v)
	}
}
// 函数类型和函数变量
func f4(a int,b int)int{
	return a+b
}
// 高阶函数
// 函数作为参数
func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)	
}
func test2(name string,f func(string)){
	f(name)
}
// 函数作为返回值
func add(a int,b int) int {
	return a+b
}
func reduce(a int,b int) int {
	return a-b	
}
func calculator(op string) func(int,int)int{
	switch op {
	   case "+":
			return add
		case "-":
			return reduce
			default:
			return nil
	}
}

