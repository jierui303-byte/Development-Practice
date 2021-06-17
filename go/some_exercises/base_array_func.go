package main

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/
import (
	"fmt"
)

func main() {
	//向函数传递数组--形参设定数组大小：
	//向函数传递数组--形参未设定数组大小：
	toFuncAndArrayData()
}

func toFuncAndArrayData() {
	/* 数组长度为 5 */
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)
	fmt.Println(avg)                    // 内部编码
	fmt.Println(float64(avg) / 1000000) // 显示

}

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
