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

	variable_function() //变量
	const_function()    //常量

}

//变量
func variable_function() int {

	//局部变量声明方式
	b := 33                  //局部变量声明
	var demoInt32 int32 = 98 //局部多变量声明
	var (
		v1 int
		v2 string
	) //局部多变量var声明

	//字符串打印
	//可以配合使用 Println()和 Printf()来打印各种自己感兴趣的信息
	fmt.Println("!... Hello World ...!")
	println(demoInt32, a, b)
	//Printf()
	fmt.Printf("字符串打印结果：%d \n", b)

	//局部变量赋值
	v1 = 12
	v2 = "hhhshshs"
	fmt.Println(v1, v2)

	return 1
}

//常量
func const_function() (ret int, err error) {
	//常量单个定义
	const APPLE_TYPE_0 int = iota
	const APPLE_TYPE_1 int = iota
	fmt.Println("常量单个定义", APPLE_TYPE_0, APPLE_TYPE_1)

	const (
		c0 = iota
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println("常量批量定义", c0, c1, c2, c3)

	const (
		size int64 = 1024
		eof        = -1 // 无类型整型常量
	)
	fmt.Println("无类型整形常量定义", size, eof)

	//常量的多重赋值
	const u, v float32 = 0, 3
	const aa, bb, cc = 3, 4, "foo"
	fmt.Println("常量多重赋值定义", u, v, aa, bb, cc)

	return 1, nil
}

//数组array 函数
func array_function() (ret int, err error) {

	//数组变量声明
	var arr1 = []int{1, 2, 3, 5, 6, 7, 8, 9}

	// Go 语言内置的函数 len()来取字符串的长度
	var v3 int = len(arr1)
	fmt.Println(arr1, v3)

	return 1, nil
}

//切片slice 函数
func slice_function() (ret int, err error) {

	return 1, nil
}

//字典map 函数
func map_function() (ret int, err error) {

	return 1, nil
}

//控制语句示例函数
func control_loop() (ret int, err error) {

	return 1, nil
}

//函数参数传递
