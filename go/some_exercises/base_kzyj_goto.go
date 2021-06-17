package main

import "fmt"

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

func main() {

	test() //goto测试代码

	//打印九九乘法表
	print9x() //嵌套for循环打印
	gotoTag() //for循环配合goto打印

}

func test() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 : 相当于把15的循环打印没有执行就直接跳到了循环开头继续循环*/
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

//嵌套for循环打印九九乘法表
func print9x() {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}
}

//for循环配合goto打印九九乘法表
func gotoTag() {
	for m := 1; m < 10; m++ {
		n := 1
	LOOP:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}
