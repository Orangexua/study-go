package main

import "fmt"

// func main() {
// 	var string = "abc"
// 	fmt.Println("hello word", string)
// }

// 空白标识符在函数返回值时的使用：

func main() {
	_,_, strs := number()
	fmt.Println(strs)
}

func number() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a,b,c
}