package main

import "fmt"

// func main() {
// 	var balance = [5]float32{1.2,2.3,4,23,23.3233242,}
// 	var balanceOne = [...]float32{1,2,3,4,5,3,2,3,4,5,3,2,0.3412345,234.5234,234.5234}
// 	fmt.Print(balance, balanceOne)
// }

func main() {
	animals := [][]string{}

	row1 := []string {"fish","shark","eel"}

	row2 := []string {"brid"}

	row3 := []string{"lizard","salamander"}

	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}