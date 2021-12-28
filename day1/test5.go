package main

import "unsafe"

// func main () {
// 	const LENGTH int = 10
// 	const WIDTH int = 5
// 	var area int
// 	const a, b, c = 1, false, "str"

// 	area = LENGTH * WIDTH
// 	fmt.Printf("面积为： %d", area)
// 	println()
// 	println(a,b,c)
// }

//常量还可以做枚举

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a,b,c)
}