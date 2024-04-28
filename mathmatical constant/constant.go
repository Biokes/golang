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
	total := 0.0
	for count := 1; count < int(number); count++ {
		total += 1.0 / factorial(float64(count))
	}
	return total
}
func factorial(number float64) float64 {
	var value = 1.0
	for count := 1; count < int(number); count++ {
		value *= number
	}
	return value
}
