package main
import "fmt"

func main() {
	// 切片
	test()
}
func test()  {
	var s1 []int
	var s2 = []int{1, 2, 3}
	var s3 =make([]int, 5)
	s3=[]int{1, 2, 3,4,5,6}
	// 长度是5，容量是10
	var s4 =make([]int, 5, 10)
	fmt.Println(cap(s1), cap(s2), cap(s3), cap(s4))// 0 3 5 10
	// 切片的赋值和传参
	s2[0] = 100
	s2[1] = 200
	// 要添加元素必须使用append()函数
    // append()会自动处理容量不足时的扩容
    // 必须接收append()的返回值（因为可能返回新切片）
    s2 = append(s2, 300) 
	fmt.Println(s2)// [100 200 3]
	fmt.Printf("len:%d cap:%d\n",cap(s2), len(s2))// 5 5
	s5 :=s3[1:3] // 从索引1开始到索引3（不包括3），左闭右开
	fmt.Println(s5)// [2 3]
	fmt.Printf("len:%d cap:%d\n",cap(s5), len(s5))// 4 4
	s6 :=s3[:3] // 从索引0开始到索引3（不包括3），左闭右开
	fmt.Println(s6)// [1 2 3]
	// 切片的复制
	// s7 :=s3[:]
	// s7[0] = 1000    // 修改s7的第一个元素，s3的第一个元素也会被修改==浅拷贝
	// fmt.Println(s3) // [1000 2 3 4 5]
	// fmt.Println(s7) // [1000 2 3 4 5]
	// 深拷贝
  s8 :=make([]int, len(s3))
	copy(s8, s3)
	s8[0] = 1000    // 修改s8的第一个元素，s3的第一个元素不会被修改
	// fmt.Println(s3) // [1 2 3 4 5]
	// fmt.Println(s8) // [1000 2 3 4 5]
	// 切片的删除
	// s = append(s[:index], s[index+1:]...)
	s8 = append(s8[:1], s8[2:]...) // 删除索引为1的元素，即第二个元素
	fmt.Println(s8) // [1000 3 4 5]
	

}