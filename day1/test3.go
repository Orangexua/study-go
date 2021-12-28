package main

import (
	"fmt"
)

// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
// var v_name v_type
// v_name = value
// func main () {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s);
// }

// 第二种，根据值自行判定变量类型。
// var v_name = value
// func main () {
// 	var d = true
// 	fmt.Println(d)
// }

func main () {
	f := "chenzixu"
	c := false
	fmt.Println(f,c)
}