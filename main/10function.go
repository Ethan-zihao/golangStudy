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
	// f3(arr)
	test(slice)
	test(arr)
	test(m)
}
// // 统一遍历函数 ，参数可能是数组、切片、map、结构体、接口
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
// 省略参数
// func f3(args ...int){
// 	for _,v:=range args{
// 		fmt.Println(v)
// 	}
// }