package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	getSquareRoot := func( x float64) float64 {
// 		return math.Sqrt(x)
// 	}

// 	fmt.Println(getSquareRoot(9))
// }

// import "fmt"

// type fun func(int) int

// func main() {
// 	testCallBack(1, callback)
// 	testCallBack(2 , func(x int) int {
// 		fmt.Printf("我是回调, x : %d \n",x)
// 		return x
// 	})
// }

// func testCallBack( x int , f fun) {
// 	f(x)
// }

// func callback( x int ) int {
// 	fmt.Printf("我是回调, x : %d \n",x)
// 	return x
// }

import "fmt"
func getSequence() func() int {
	i:=0
	return func() int {
	   i+=1
	  return i  
	}
 }
 
 func main(){
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()  
 
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	
	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()  
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
 }