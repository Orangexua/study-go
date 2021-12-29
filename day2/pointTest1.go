package main

import "fmt"

func main() {
	var ptr *int
	n := (ptr == nil)
	fmt.Printf("ptr 的值为: %x \n", ptr)
	fmt.Print(n)
}