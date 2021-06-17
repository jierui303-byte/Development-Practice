package main

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

func main() {

	if1()
	if2()
	if3()

}

func if1() {
	var number int = 0

	//if-else
	if 100 < number {
		number++
	} else {
		number--
	}
}

func if2() {
	var number int = 0

	//因为if语句可以接受一条初始化子语句，所以我们常常会使用它来初始化一个变量：
	if diff := 100 - number; 100 < diff {
		number++
	} else {
		number--
	}
}

func if3() {
	var number int = 0

	//if-elseif-else [初始化子句和条件表达式之间是需要用分号“;”分隔的]
	if diff := 100 - number; 100 < diff {
		number++
	} else if 200 < diff {
		number--
	} else {
		number -= 2
	}
}
