package main

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

import (
	"fmt"
)

func main() {

	//常量单个定义
	const APPLE_TYPE_0 int = iota
	const APPLE_TYPE_1 int = iota
	fmt.Println("常量单个定义", APPLE_TYPE_0, APPLE_TYPE_1)

	/**********************************************************************************************************/

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

	/**********************************************************************************************************/

	//常量的多重赋值
	const u, v float32 = 0, 3
	const aa, bb, cc = 3, 4, "foo"
	fmt.Println("常量多重赋值定义", u, v, aa, bb, cc)

}
