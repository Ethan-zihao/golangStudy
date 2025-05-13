package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// go不能用0和非0来表示true和false
	var b bool = true
	fmt.Printf("b: %v\n", b)
	// g
	// go语言中数据类型不能隐式转换，必须显示转换
	// go语言中数据类型包括: 基本数据类型，数组，切片，结构体，函数，接口，通道，映射，指针
	// 基本数据类型包括: 整型，浮点型，布尔型，字符串，字符型，复数型
	// 整型包括: int8，int16，int32，int64，uint8，uint16，uint32，uint64，int，uint，uintptr
	// int8(-128~127)，int16(-32768~32767)，int32(-2147483648~2147483647)，int64(-9223372036854775808~9223372036854775807)
	// uint8(0~255)，uint16(0~65535)，uint32(0~4294967295)，uint64(0~18446744073709551615)
	// 浮点型包括: float32，float64 ，complex64，complex128
	// 布尔型包括: bool
	// 字符串包括: string
	// 字符型包括: byte，rune，string，byte和rune都是uint8类型，string是一个字符串类型
	// go语言中字符串是不可变的，字符串是一个字节数组，字符串是一个只读的字节数组，字符串不能被修改
	// 但是字符串可以被拼接，字符串可以被切割，字符串可以被转换为其他类型，字符串可以被比较，字符串可以被遍历，字符串可以被格式化
	name := "John"
	age := 30
	// golang中字符串是不可变的,每一次字符串的拼接都会创建一个新的字符串,所以字符串拼接的效率很低
	// (1)使用fmt.Sprintf函数可以提高字符串拼接的效率，fmt.Sprintf函数的返回值是一个字符串，fmt.Sprintf函数的第一个参数是格式化字符串，
	// 后面的参数是格式化字符串中的占位符，占位符的类型和数量必须与格式化字符串中的占位符的类型和数量一致
	// msg := name + " is " + fmt.Sprint(age) + " years old."
	msg := fmt.Sprintf("%s is %d years old.", name, age)
	// %s表示字符串占位符，%d表示整数占位符，%f表示浮点数占位符，%t表示布尔值占位符，%v表示值占位符，%T表示类型占位符
	fmt.Printf("name: %v\n", msg)
	//   (2)string.Join()
	// string.Join()函数的第一个参数是分隔符，第二个参数是字符串数组，string.Join()函数的返回值是一个字符串
	// join会先根据分隔符将字符串数组内容，计算出一个字符串的长度，然后再根据字符串的长度，分配一个字符串的内存空间，
	msg2 := strings.Join([]string{name, fmt.Sprint(age)}, " is ")
	fmt.Printf("name: %v\n", msg2)
	// (3)string.Builder()
	// string.Builder()函数的返回值是一个字符串构建器，字符串构建器是一个可变的字符串，字符串构建器可以被拼接，字符串构建器可以被转换为字符串
	msg3 := strings.Builder{}
	msg3.WriteString(name)
	msg3.WriteString(" is ")
	msg3.WriteString(fmt.Sprint(age))
	fmt.Printf("name: %v\n", msg3.String())
	// (4)buffer.writeString()
	// buffer.WriteString()函数的返回值是一个字符串构建器，字符串构建器是一个可变的字符串，字符串构建器可以被拼接，字符串构建器可以被转换为字符串
	var buffer bytes.Buffer
	buffer.WriteString(name)
	buffer.WriteString(" is2 ")
	buffer.WriteString(fmt.Sprint(age))
	fmt.Printf("name: %v\n", buffer.String())
	// (5)字符串切片
	// 字符串切片是一个字符串数组，字符串切片可以被拼接，字符串切片可以被切割，字符串切片可以被转换为字符串，字符串切片可以被比较，字符串切片可以被遍历，字符串切片可以被格式化
	// 字符串切片的语法为: []string{a, b, c}，a，b，c为字符串切片的元素，字符串切片的元素类型为string
	fmt.Printf("name: %v\n", name[0:2])
	// (6)字符串常用方法
	// 字符串常用方法包括: 字符串长度，字符串拼接，字符串切割，字符串转换为其他类型，字符串比较，字符串遍历，字符串格式化
	// 1.len()
	// nameLen := len(name)
	// 2.strings.contains()
	// fmt.Printf("name: %v\n", strings.Contains(name, "John"))
	// 3.strings.Split()
	fmt.Printf("name: %v\n", strings.Split(name, "John"))
	// 4.strings.ToUpper()
	fmt.Printf("name: %v\n", strings.ToUpper(name))
	// 5.strings.ToLower()
	fmt.Printf("name: %v\n", strings.ToLower(name))
	// 6.strings.Replace() // strings.Replace()函数的返回值是一个字符串，strings.Replace()函数的第一个参数是源字符串，第二个参数是替换字符串，第三个参数是替换次数，-1表示全部替换
	fmt.Printf("name: %v\n", strings.Replace(name, "John", "Tom", -1))
	// 7.stings.Index() // strings.Index()函数的返回值是一个整数，strings.Index()函数的第一个参数是源字符串，第二个参数是查找字符串
	fmt.Printf("name: %v\n", strings.Index(name, "n"))
	// 8.strings.LastIndex() // strings.LastIndex()函数的返回值是一个整数，strings.LastIndex()函数的第一个参数是源字符串，第二个参数是查找字符串
	fmt.Printf("name: %v\n", strings.LastIndex(name, "n"))
	// 9.strings.Trim() // strings.Trim()函数的返回值是一个字符串，strings.Trim()函数的第一个参数是源字符串，第二个参数是去除的字符串
	fmt.Printf("name: %v\n", strings.Trim(name, "John"))
	// 复数型包括: complex64，complex128
	// 数组包括: [n]T，n为数组长度，T为数组元素类型
	// 切片包括: []T，T为切片元素类型
	// 结构体包括: struct{}，struct{name string, age int}
	// 函数包括: func()，func(a int) int，func(a int, b int) (int, int)
	// 接口包括: interface{}，interface{name() string}
	// 通道包括: chan T，T为通道元素类型
	// 映射包括: map[K]V，K为键类型，V为值类型
	// 指针包括: *T，T为指针指向的类型
	// 类型转换包括: T(a)，T为目标类型，a为源类型
	// 类型断言包括: a.(T)，a为源类型，T为目标类型

	var a int = 10
	var c float64 = float64(a)
	fmt.Printf("c: %v\n", c)
	// 类型转换
	// 类型转换是将一个类型的值转换为另一个类型的值，类型转换的语法为: T(a)，T为目标类型，a为源类
	var a1 int = 10
	// 将a1转换为float64类型
	var c1 float64 = float64(a1)
	fmt.Printf("c1: %vv\n", c1)
	// 类型断言
	// 类型断言是将一个接口类型的值转换为另一个类型的值，类型断言的语法为: a.(T)，a为源类型，T为目标类型
	var a2 interface {
		// 定义一个接口类型
		// 接口类型可以包含多个方法，方法可以有参数和返回值，方法可以有名称，方法可以有实现，方法可以有实现体
	} = 10
	// 将a2转换为int类型
	var c2 int = a2.(int)
	fmt.Printf("c2: %v\n", c2)
}
