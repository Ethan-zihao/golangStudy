package main
import (
	"fmt"	

)
func main() {
	// golang的初始化顺序：
	// 1. 变量初始化
	// 2. 常量初始化
	// 3. 导入包
	// 4. 初始化函数
	// 5. main函数
	// 6. 其他函数
	// golang有一个特殊的函数init函数，先于main函数执行，实现包级别的初始化操作。
	// init函数的作用是在包被导入时执行，用于初始化包级别的变量、常量、函数等。
	//init主要特点
	// 1. 函数名必须是init
	// 2. 函数不能有参数和返回值
	// 3. 函数不能被其他函数调用
	// 4. 函数不能被其他包调用
	// 每个包可以有多个init函数，它们会按照它们在代码中的出现顺序依次执行。
	// 每个包的init函数只会执行一次，即使包被多次导入。
	// 每个包的init函数会在main函数之前执行。
	// 不同包的init函数会按照它们在代码中的出现顺序依次执行。
	
  fmt.Println("main...")



}
var i int = initVar()
var j =5
    func init (){  //finit
		fmt.Println("init")
       j=10
       fmt.Println(j)
	}
   func init (){
	   fmt.Println("init2")
   }
  
    func initVar () int{
		fmt.Println("initVar...")
        return 100
	}