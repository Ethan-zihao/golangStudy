package main
import "fmt"

func main() {
    // arr := [5]int{1,2,3,4,5}
    // f2(arr[:])
	// var m map[string]int      // 声明map
    //  m := make(map[string]int) // 使用make创建
	n := map[string]int{"a":1, "b":2}
	// 通道返回值
	// map返回键和值
	
     f4(n)
}

// func f2(slice []int) {
//     for _,v := range slice {
//         fmt.Printf("%d ", v)
//     }
//     fmt.Println()
// }
// 遍历map对象
func f4 (m map[string]int){
		for k,v := range m {
			fmt.Printf("%s:%d", k,v)
}}