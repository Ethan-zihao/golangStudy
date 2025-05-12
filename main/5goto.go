package main
import "fmt"

func main() {
    // for range遍历数组、切片、字符串、通道、map
    // 数组、切片、字符串返回索引和值

   // 	数组(Array)
// 固定长度：声明时必须指定长度，且长度不可变
// 值类型：赋值和传参会复制整个数组，而不是指针
arr := [5]int{1,2,3,4,5}  // 声明长度为5的int数组
// arr := [3]string{"a", "b", "c"} // 声明并初始化
// f2(arr[:])
//    切片(Slice)
// 动态长度：长度可变，有容量(capacity)概念
// 引用类型：底层引用一个数组，赋值和传参传递的是引用
// var s []int               // 声明切片
// s := make([]int, 5)       // 使用make创建
// s := []int{1, 2, 3}       // 声明并初始化
// s := arr[1:3]             // 从数组创建切片
//    // 映射(Map)
// // 键值对集合：无序的key-value集合
// // 引用类型：必须使用make初始化后才能使用
var m map[string]int      // 声明map
m := make(map[string]int) // 使用make创建
m := map[string]int{"a":1, "b":2} // 声明并初始化
	// 通道返回值
	// map返回键和值
	
f4(m)


}

	// 遍历数组、切片、字符串

// func f2(slice []int) {  // 改为接收切片参数
//     for _,v := range slice {
//         fmt.Printf("%d ", v)
//     }
//     fmt.Println()
// }
		// 遍历切片
	// func f3(Slice){
	//   for _,v :=range slice {
	//    fmt.Printf("%d", v)  
	//  }
	// }
	// 遍历map
	func f4 (map [string]int){
		for k,v := range map {
			fmt.Printf("%s:%d", k,v)
	}
}