package main

import "fmt"

func main() {
	value := 0
	for value < 0 {
		fmt.Print("Enter your number(greater than 0) to get the factorial: ")
		_, _ = fmt.Scanln(&value)
	}
	fmt.Printf("Factorial of %d is %d\n", value, getFactorial(value))
}
func getFactorial(value int) int {
	number := 1
	for value != 0 {
		number *= value
		value--
	}
	return number
}
