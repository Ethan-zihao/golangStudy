package main
import (
	"fmt"
	"os"

)
func main() {
	// go语言中defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
	// defer语句的语法是：defer 函数名(参数)。
	// defer特性：
	// 1. 当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行函数下一个语句。
	// 2. 当函数执行完毕后，再从defer栈中，依次从栈顶取出语句执行(注：遵守栈先入后出的机制)。
	// 3. defer语句可以读取函数中的命名返回值。
	// 4. defer语句的执行顺序与它们在代码中的出现顺序相反。
	// 5. defer语句可以用于关闭文件、释放资源等操作。
    // 实例：
    // defer语句的使用
    defer fmt.Println("defer1")
    defer fmt.Println("defer2")
    defer fmt.Println("defer3")

}
func readFile() {
   file,_ :=os.open("1var.go")
   defer file.Close()  // 确保文件一定会被关闭
    
    // 文件操作...
    // 即使中间有return或panic，defer也会执行
    // 即使函数执行完毕，defer也会执行
}