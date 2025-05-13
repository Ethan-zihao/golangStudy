package main

import "fmt"

func test() {
	// 切片遍历
	var s1 = []int{1, 2, 4, 6, 6}

	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("s1[%v]: %v\n", i, s1[i])
	// }
	// for i, v := range s1 {
	// 	fmt.Printf("i: %v\n", i)
	// 	fmt.Printf("v: %v\n", v)
	// }

	// 切片CRUD

	// 在  append  方法中，第一个参数是要追加元素的切片，第二个参数是要追加的元素。具体解释如下：
	// 末尾增加元素
	// 在索引为2的位置插入元素6
	// insertIndex := 2
	// insertValue := 6

	// 将切片划分为两部分，然后拼接插入元素
	// slice = append(slice[:insertIndex], append([]int{insertValue}, slice[insertIndex:]...)...)

	// fmt.Println(slice)

	// s1 = append(s1[:2], append([]int{99}, s1[2:]...)...)
	// fmt.Printf("s1: %v\n", s1)
	// 在上述示例中，我们使用  append  函数将切片划分为两部分。首先，我们使用  slice[:insertIndex]  获取索引为  insertIndex  之前的部分，然后使用  slice[insertIndex:]  获取索引为  insertIndex  之后的部分。接下来，我们使用  append  函数将  insertValue  插入到两个切片之间，形成最终的切片。

	// 请注意，这种方法会创建一个新的切片，而不是在原始切片上进行插入操作。如果需要在原始切片上进行插入操作，可以使用  copy  函数将元素后移，然后将新元素插入到指定位置。

	// 扩展切片长度
	// 	 slice = append(slice, 0)

	// 	 // 将插入位置之后的元素后移一位
	// 	 copy(slice[insertIndex+1:], slice[insertIndex:])

	// 	 // 在指定位置插入新元素
	// 	 slice[insertIndex] = insertValue
	// 	 在上述示例中，我们首先使用  append  函数将切片长度扩展一个位置，然后使用  copy  函数将插入位置之后的元素后移一位。最后，我们将新元素  insertValue  插入到指定位置  insertIndex 。

	// 这种方法会在原始切片上进行插入操作，而不会创建新的切片。但请注意，如果原始切片已经达到了容量限制，这种方法可能会导致运行时错误。
	s1 = append(s1, 100)
	fmt.Printf("s1: %v\n", s1)
	copy(s1[3:], s1[2:])
	s1[2] = 98
	fmt.Printf("s1: %v\n", s1)

	// 删除元素  公式
	// append(a[:index],a[index+:]...)
	// s1 = append(s1[:2], s1[3:]...)
	// fmt.Printf("s1: %v\n", s1)
	// indexkey := 100
	// for i, v := range s1 {
	// 	if indexkey == v {
	// 		fmt.Printf("i: %v\n", i)
	// 	}
	// }

	//切片拷贝
	// var s2 []int
	var s2 = make([]int, 8)
	copy(s2, s1)
	fmt.Printf("s2: %v\n", s2)

}

func main() {
	test()
}
