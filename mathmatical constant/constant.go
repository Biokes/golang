package main

import "fmt"

func main() {
	fmt.Println(collectInput())
}
func collectInput() float64 {
	number := 0
	for number < 1 {
		fmt.Println("Enter a number to get its maths value!!: ")
		number, _ = fmt.Scanln(&number)
	}
	return getConstant(number)
}
func getConstant(number int) float64 {
	value := 1.0
	counter := 0
	total := 0.0
	for counter <= number {
		total += 1.0 / factorial(value)
		counter++
	}
	return total
}
func factorial(number float64) float64 {
	var value int = 1
	for number != 1 {
		value *= number
		number -= 1.0
	}
	return number
}
