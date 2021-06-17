package main

import (
	"fmt"
	"time"
)

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

func main() {
	//createGoroutine()            //协程创建使用
	//createChannel()              //通道创建使用
	//createChannelCacheSize()     //设定通道缓冲区大小buffer
	loopChannelAndCloseChannel() //Go 遍历通道与关闭通道

	channelAndSelectCase()      //channel结合select和case
	channelSelectTest()         //select的随机性	测试
	chanReadOnlyOrReceiveOnly() //channel支持3种类型（通过%T看到的）
}

//创建一个协程来执行
func createGoroutine() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//创建一个通道缓冲区
func createChannel() {
	/***
		make(chan string, 2)
		表示可以容纳2个string的内容，比如"abcdefg"和"h"这样2个string，而不是指只能存"a"和"b"。
		当不设置length时候即为buffer=0
	***/

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) //创建通道，第二个buffer不设置时默认为0，代表容纳多少个int的内容数量，而非具体的内存容积
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

//通道缓冲区--通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：ch := make(chan int, 100)
func createChannelCacheSize() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据【不设定时默认只能发送一条数据】
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func loopChannelAndCloseChannel() {
	//Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：
	//	v, ok := <-ch
	//如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}

	//len()可以用来查看数组或slice的长度
	//cap()可以用来查看数组或slice的容量
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //关闭管道【只有发送者才可以关闭管道】
}

//channel结合select和case
func channelAndSelectCase() {
	/***
	select和case的组合可以使哪个管道就绪（对端已阻塞），就读取该管道数据并执行相应case的代码块。

	官网译: select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。
		***/
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	quit := make(chan int)
	go receive(ch1, ch2, ch3, quit)
	send(ch1, ch2, ch3, quit)
}

func receive(ch1, ch2, ch3, quit chan int) {
	for i := 0; i < 2; i++ {
		fmt.Printf("receive %d from ch1\n", <-ch1)
		fmt.Printf("receive %d from ch2\n", <-ch2)
		fmt.Printf("receive %d from ch3\n", <-ch3)
	}
	quit <- 0
}

func send(ch1, ch2, ch3, quit chan int) {
	//select支持default
	//当所有case都阻塞的时候，执行default
	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:
			fmt.Printf("send %d to ch1\n", i)
		case ch2 <- i:
			fmt.Printf("send %d to ch2\n", i)
		case ch3 <- i:
			fmt.Printf("send %d to ch3\n", i)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//select的随机性 测试
func channelSelectTest() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go receiveTest(ch1)
	go receiveTest(ch2)
	go receiveTest(ch3)
	sendTest(ch1, ch2, ch3)
}

func receiveTest(ch chan int) {
	for {
		<-ch
	}
}

func sendTest(ch1, ch2, ch3 chan int) {
	for i := 0; i < 10; i++ {
		// sleep是为了保证所有的管道receiver都已阻塞等待数据
		time.Sleep(1000 * time.Millisecond)
		select {
		case ch1 <- i:
			fmt.Printf("send %d to ch1\n", i)
		case ch2 <- i:
			fmt.Printf("send %d to ch2\n", i)
		case ch3 <- i:
			fmt.Printf("send %d to ch3\n", i)
		}
	}
}

//channel支持3种类型使用注意事项
func chanReadOnlyOrReceiveOnly() {
	/***
	channel支持3种类型（通过%T看到的）：

	读写类型：chan int

	只读类型：<-chan int，叫做receive-only

	只写类型：chan<- int，叫做send-only

	通过函数参数传递时候，若原来是读写，则可以转成只读或只写，但如果已经是只读或只写，则只能保持类型，无法转为其他类型（比如原来是只读，则只能是只读，无法转成只写，也无法转为读写），例子如下：

	***/
	c := make(chan int)
	fmt.Printf("%T\n", c)
	go sendz(c)
	go recvz(c)
	time.Sleep(1 * time.Second)
}
func sendz(c chan<- int) {
	fmt.Printf("send: %T\n", c)
	c <- 1
}

func recvz(c <-chan int) {
	fmt.Printf("recv: %T\n", c)
	fmt.Println(<-c)
}
