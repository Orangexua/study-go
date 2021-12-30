package main

import "fmt"

func main() {
	number := []int {1,2,3,4,5,6,7,8,9,10}

	printSlice(number)
}

func printSlice(x []int) {
	fmt.Printf("len%d cap=%d slice=%v\n", len(x), cap(x),x)
}