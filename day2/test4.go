package main

import "fmt"

func main() {
	num1 := 2
	num2 := 10
	ret := max(num1, num2)

	fmt.Printf("最大值是: %d\n", ret)
}

func max( num1 , num2 int) int {
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}