package main

import "fmt"

func main() {
	fmt.Println("task1 =====================================")
	input := []int{1, 2, 3, 4, 5}
	sumResult := SumNumbers(input)
	fmt.Println(sumResult)
	fmt.Println("")

	manageProduct()
	manageOrders()
}
