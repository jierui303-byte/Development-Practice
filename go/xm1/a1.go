package main

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

*/

import (
	"fmt"
)

//全局变量声明
var a int = 10

func main() {

	//字符串打印 可以配合使用 Println()和 Printf()来打印各种自己感兴趣的信息
	fmt.Println("!... Hello World ...!")

	//局部变量声明
	b := 33

	//局部多变量声明
	var demoInt32 int32 = 98
	println(demoInt32, a, b)

	//局部多变量var声明
	var (
		v1 int
		v2 string
		v3 int
	)
	//局部变量赋值
	v1 = 12
	v2 = "hhhshshs"

	//数组变量声明
	var arr1 = []int{1, 2, 3, 5, 6, 7, 8, 9}
	// Go 语言内置的函数 len()来取字符串的长度
	v3 = len(arr1)
	fmt.Println(v1, v2, v3)

}
