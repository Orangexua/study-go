package main

import "fmt"
var a int = 20

func main() {
	var a int = 10 
	var b int = 20 
	var c int = 0

	fmt.Printf("main()函数a = %d \n",a)
	c = sum(a, b)
	fmt.Printf("main()函数c = %d \n",c)
}

func sum (a, b int) int {
	fmt.Printf("sum() 函数 a = %d\n", a)
	fmt.Printf("sum() 函数 b = %d\n", b)
	return a + b
}

