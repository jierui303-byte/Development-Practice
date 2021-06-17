package xm1_base

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

import (
	"fmt"
)

//全局变量声明
var a int = 10

//定义一个vInterface接口，里面包含一个speak()方法
type vInterface interface {
	Speak() string
}

func main() {

	variable_function() //变量
	const_function()    //常量

	var v interface{}
	control_loop(v)

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

	/**********************************************************************************************************/

	//字符串打印
	//可以配合使用 Println()和 Printf()来打印各种自己感兴趣的信息
	fmt.Println("!... Hello World ...!")
	println(demoInt32, a, b)
	//Printf()
	fmt.Printf("字符串打印结果：%d \n", b)

	/**********************************************************************************************************/

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

func getContent() (ret string, err error) {
	return "Python", nil
}

//控制语句示例函数
func control_loop(v interface{}) (ret int, err error) {

	var number int = 0

	//if-else
	if 100 < number {
		number++
	} else {
		number--
	}

	//因为if语句可以接受一条初始化子语句，所以我们常常会使用它来初始化一个变量：
	if diff := 100 - number; 100 < diff {
		number++
	} else {
		number--
	}

	//if-elseif-else [初始化子句和条件表达式之间是需要用分号“;”分隔的]
	if diff := 100 - number; 100 < diff {
		number++
	} else if 200 < diff {
		number--
	} else {
		number -= 2
	}

	/**********************************************************************************************************/

	//for循环  number是一个int类型的变量
	for number < 200 { // 当number大于等于200时for循环会退出。
		number += 2
	}

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

	//range子句
	ints := []int{1, 2, 3, 4, 5}
	for i, d := range ints {
		fmt.Printf("%d: %d\n", i, d)
	}

	var i, d int
	for i, d = range ints {
		fmt.Printf("%d: %d\n", i, d)
	}

	ints := []int{1, 2, 3, 4, 5}
	length := len(ints)
	indexesMirror := make([]int, length)
	elementsMirror := make([]int, length)
	var i int
	for indexesMirror[length-i-1], elementsMirror[length-i-1] = range ints {
		i++
	}

	ints := []int{1, 2, 3, 4, 5}
	for i := range ints {
		d := ints[i]
		fmt.Printf("%d: %d\n", i, d)
	}

	m := map[uint]string{1: "A", 6: "C", 7: "B"}
	var maxKey uint
	for k := range m {
		if k > maxKey {
			maxKey = k
		}
	}

	var values []string
	for _, v := range m {
		values = append(values, v)
	}

	/**********************************************************************************************************/

	//goto语句
	/***
		goto使用注意事项：
			第一，不允许因使用goto语句而使任何本不在当前作用域中的变量进入该作用域
			第二，我们把某条goto语句的直属代码块叫作代码块A，而把该条goto语句右边的标记所指代的那条标记语句的直属代码块叫作代码块B。
	***/
	goto L
	v := "B"
L:
	fmt.Printf("V: %v\n", v)

	/**********************************************************************************************************/

	//switch简单用法
	content, err := getContent()
	switch content {
	default:
		fmt.Println("Unknown language")
	case "Python":
		fmt.Println("A interpreted Language")
	case "Go":
		fmt.Println("A compiled language")
	}

	//switch配合fallthrough语句
	switch content, _ := getContent(); content {
	default:
		fmt.Println("Unknown language")
	case "Ruby":
		fallthrough
	case "Python":
		fmt.Println("A interpreted Language")
	case "C", "Java", "Go":
		fmt.Println("A compiled language")
	}

	//switch配合break语句
	switch content, _ := getContent(); content {
	default:
		fmt.Println("Unknown language")
	case "Ruby":
		break
	case "Python":
		fmt.Println("A interpreted Language")
	case "C", "Java", "Go":
		fmt.Println("A compiled language")
	}

	//switch的另外一种变形写法
	/**
		在类型switch语句中，case表达式中的类型字面量可以是nil。
	在前面的那个示例中，如果v的值是nil，那么表达式v.(type)的结果值也会是nil。
	因此，当这种情况发生时，如果存在包含了nil的case表达式，那么与它相对应的那个分支就会被执行
		**/
	switch i := v.(type) {
	case string:
		i := v.(string)
		fmt.Printf("The string is '%s'.\n", i)
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		fmt.Printf("The integer is %d.\n", i)
	default:
		fmt.Printf("Unsupported value. (type=%T)\n", i)
	}

	//switch的另外一种变形写法
	//在switch表达式缺失的情况下，该switch语句的判定目标会被视为布尔类型值true.
	//也就是说，其中的所有case表达式的结果值都应该是布尔类型的。
	//并且，自上而下，第一个结果值为true的case表达式所对应的分支会被执行。
	switch {
	case number < 100:
		number++
	case number < 200:
		number--
	default:
		number -= 2
	}

	//switch的另外一种变形写法
	switch number := 123; {
	case number < 100:
		number++
	case number < 200:
		number--
	default:
		number -= 2
	}

	/**********************************************************************************************************/

	//select

	//break,fallthrough,

	return 1, nil
}

//函数参数传递[指定参数个数传递]
func function_prarams_function() (ret int, err error) {

	return 1, nil
}

//函数参数传递[不定参数个数传递]
func function_prarams_args_function() (ret int, err error) {

	return 1, nil
}
