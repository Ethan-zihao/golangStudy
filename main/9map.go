package main
import "fmt"
func main() {
	// map
	test()

}
func test (){

	var fruit=map[string]int{"apple":1,"banana":2,"orange":3}
	    // 添加元素
		fruit["pear"]=4  
		fmt.Println(fruit)
		// 获取元素
		fmt.Println(fruit["apple"])  
		// 删除元素
		delete(fruit,"apple")
		fmt.Println(fruit)
		// 检查元素是否存在
		// value, ok := map[key]
		value,ok:=fruit["banana"]
		fmt.Println(value,ok)
		// 遍历map
		for key,value:=range fruit{
			fmt.Println(key,value)
		}
}