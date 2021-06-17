package main

import "fmt"

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

/* 定义接口 */
//type interface_name interface {
//	method_name1 [return_type]
//	method_name2 [return_type]
//	method_name3 [return_type]
//	...
//	method_namen [return_type]
// }

/* 定义结构体 */
// type struct_name struct {
/* variables */
// }

/* 实现接口方法 */
// func (struct_name_variable struct_name) method_name1() [return_type] {
/* 方法实现 */
// }
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
/* 方法实现*/
//}

//定义一个接口	Phone
type Phone interface {
	call() //接口里有一个方法叫call()
}

//定义一个结构体 NokiaPhone
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

//定义一个结构体 IPhone
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	createInterfaceAndStruct() //结构体和接口配合

	//理解Golang中的interface和interface{}:https://www.cnblogs.com/maji233/p/11178413.html
}

func createInterfaceAndStruct() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
