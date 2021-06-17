package main

import "fmt"

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

func main() {
	createSlice()
	emptySlice()
	cutSlice()
	addContentSlice()
}

//切片定义方式
func createSlice() {
	//定义切片方式一：
	//var slice1 []type = make([]type, len)

	//定义切片方式一：
	//slice1 := make([]type, len)
	//slice1 := make([]T, length, capacity)	//也可以指定容量，其中 capacity 为可选参数。这里 len 是数组的长度并且也是切片的初始长度。

	//切片初始化
	var numbers = make([]int, 3, 5)

	//切片是可索引的，并且可以由 len() 方法获取长度。
	//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
	printSlice(numbers)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

//新建空切片
func emptySlice() {
	//创建一个空切片(nil)
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

//切片截取
func cutSlice() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

//切片容量增加
//如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
//下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
func addContentSlice() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
