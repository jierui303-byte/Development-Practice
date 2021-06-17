package main

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

import (
	"fmt"
)

func getContent() (ret string, err error) {
	return "Python", nil
}

func main() {

	switch1()
	switch2()
	switch3()
	switch4()
	switch5()
	switch6()

}

func switch1() {
	//switch简单用法
	content, _ := getContent()
	switch content {
	default:
		fmt.Println("Unknown language")
	case "Python":
		fmt.Println("A interpreted Language")
	case "Go":
		fmt.Println("A compiled language")
	}
}

func switch2() {
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
}

func switch3() {
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
}

func switch4() {
	/**
	switch的另外一种变形写法
		在类型switch语句中，case表达式中的类型字面量可以是nil。
	在前面的那个示例中，如果v的值是nil，那么表达式v.(type)的结果值也会是nil。
	因此，当这种情况发生时，如果存在包含了nil的case表达式，那么与它相对应的那个分支就会被执行
		**/
	var v interface{}

	switch i := v.(type) {
	case string:
		//i := v.(string)
		fmt.Printf("The string is '%s'.\n", i)
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		fmt.Printf("The integer is %d.\n", i)
	default:
		fmt.Printf("Unsupported value. (type=%T)\n", i)
	}
}

func switch5() {
	//switch的另外一种变形写法
	//在switch表达式缺失的情况下，该switch语句的判定目标会被视为布尔类型值true.
	//也就是说，其中的所有case表达式的结果值都应该是布尔类型的。
	//并且，自上而下，第一个结果值为true的case表达式所对应的分支会被执行。
	var number int = 0

	switch {
	case number < 100:
		number++
	case number < 200:
		number--
	default:
		number -= 2
	}
}

func switch6() {
	//switch的另外一种变形写法
	switch number := 123; {
	case number < 100:
		number++
	case number < 200:
		number--
	default:
		number -= 2
	}
}
