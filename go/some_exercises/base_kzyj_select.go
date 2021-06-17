package main

import "fmt"

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

func main() {

	/***
	Channel是Go中的一个核心类型，你可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯(communication)。

	它的操作符是箭头 <- 。

	ch <- v    // 发送值v到Channel ch中
	v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v
	(箭头的指向就是数据的流向)

	就像 map 和 slice 数据类型一样, channel必须先创建再使用:

	ch := make(chan int)

	Channel类型的定义格式如下：

	ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
	它包括三种类型的定义。可选的<-代表channel的方向。如果没有指定方向，那么Channel就是双向的，既可以接收数据，也可以发送数据。

	chan T          // 可以接收和发送类型为 T 的数据
	chan<- float64  // 只可以用来发送 float64 类型的数据
	<-chan int      // 只可以用来接收 int 类型的数据
	<-总是优先和最左边的类型结合。

	chan<- chan int    // 等价 chan<- (chan int)
	chan<- <-chan int  // 等价 chan<- (<-chan int)
	<-chan <-chan int  // 等价 <-chan (<-chan int)
	chan (<-chan int)
	使用make初始化Channel,并且可以设置容量:

	make(chan int, 100)


	详细文档：https://www.runoob.com/w3cnote/go-channel-intro.html

	***/
	select1()

}

func select1() {
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
