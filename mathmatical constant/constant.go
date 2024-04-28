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
	total := 0.0
	var input float64 = float64(number)
	for value <= input {
		total += 1 / factorial(value)
		fmt.Println(total)
		value++
	}
	return total
}
func factorial(number float64) float64 {
	var value float64 = 1
	for number != 1 {
		value *= number
		number--
	}
	return number
}
