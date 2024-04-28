package main

import "fmt"

func main() {
	collectInput()
}
func collectInput() float64 {
	var number int
	fmt.Println("Enter a number to get its maths value: ")
	number, _ = fmt.Scanln(&number)
	for number <= 1 {
		fmt.Println("Enter a number to get its maths value!!: ")
		number, _ = fmt.Scanln(&number)
	}
	return getConstant(number)
}
func getConstant(number int) float64 {
	value := 1
	total := 0.0
	for value <= number {
		total += float64(value / factorial(value))
		value++
	}
	return total
}
func factorial(number int) int {
	value := 1
	for number != 0 {
		value *= number
		number--
	}
	return number
}
