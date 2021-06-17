package xm1_base

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

func main() {

	for1()
	for2()

}

func for1() {
	//for循环  number是一个int类型的变量
	for number < 200 { // 当number大于等于200时for循环会退出。
		number += 2
	}
}

func for2() {

	//for子句：一条for语句可以携带一个for子句，并可以使用这个for子句提供的条件来对迭代进行控制。
	for i := 0; i < 100; i++ {
		number++
	}
	var j uint = 1
	for ; j%5 != 0; j *= 3 { // 省略初始化子句
		number++
	}
	for k := 1; k%5 != 0; { // 省略后置子句
		k *= 3
		number++
	}
}
