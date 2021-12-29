package main

import "fmt"

func main() {
	s := 1
	for i := 0; i <= 10; i++ {
		s += i
	}
	fmt.Println(s)
}
