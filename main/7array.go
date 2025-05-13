package main
 import "fmt"
func main() {
	test2()
}
func test2() {
    var arr [5]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    arr[3] = 4
    arr[4] = 5
    // for i, v := range arr {  // 遍历数组，i是索引，v是值
    //     fmt.Printf("arr[%d] = %d\n", i, v)
    // }
	// 数组的初始化
   var arr2 =[...]int{1,2,3,4,5}
    // 访问数组元素
   fmt.Println(arr2[0])
   fmt.Printf("%d\n",arr2[len(arr2)-1])
    

}