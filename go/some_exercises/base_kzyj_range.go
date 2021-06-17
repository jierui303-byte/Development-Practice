package main

/*
 在go中，则是任何一个package中，都可以有唯一一个带有main方法的go文件

 import引入包有三种方式：

*/

import (
	"fmt"
)

func main() {

	range1()
	range2()
	range3()
	range4()
	range5()
	range6()

}

func range1() {
	ints := []int{1, 2, 3, 4, 5}
	for i, d := range ints {
		fmt.Printf("%d: %d\n", i, d)
	}
}

func range2() {
	var i, d int
	for i, d = range ints {
		fmt.Printf("%d: %d\n", i, d)
	}
}

func range3() {
	ints := []int{1, 2, 3, 4, 5}
	length := len(ints)
	indexesMirror := make([]int, length)
	elementsMirror := make([]int, length)
	var i int
	for indexesMirror[length-i-1], elementsMirror[length-i-1] = range ints {
		i++
	}

}

func range4() {
	ints := []int{1, 2, 3, 4, 5}
	for i := range ints {
		d := ints[i]
		fmt.Printf("%d: %d\n", i, d)
	}
}

func range5() {
	m := map[uint]string{1: "A", 6: "C", 7: "B"}
	var maxKey uint
	for k := range m {
		if k > maxKey {
			maxKey = k
		}
	}
}

func range6() {
	var values []string
	for _, v := range m {
		values = append(values, v)
	}
}
