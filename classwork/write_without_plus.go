package main

import "fmt"

func main() {
	fmt.Println("Enter number1: ")
	var (
		num1 int
		num2 int
	)
	_, _ = fmt.Scan(&num1)
	fmt.Println("Enter number2: ")
	_, _ = fmt.Scan(&num2)
	fmt.Println(add(num1, num2))
}
func add(num1 int, num2 int) int {
	return num1 - (-1 * (num2))
}
