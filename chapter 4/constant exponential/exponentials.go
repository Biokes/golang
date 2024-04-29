package main

import "fmt"

func main() {
	number := collectInput()
	fmt.Println(calculate(number))
}
func collectInput() int {
	var number int
	fmt.Println("Enter a number: ")
	_, _ = fmt.Scanln(&number)
	return number
}
func calculate(given int) float64 {
	total := float64(1)
	var numerator = given
	for count := 1; count <= given; count++ {
		numerator *= numerator
		total += float64(numerator / factorial(count))
	}
	return total
}
func factorial(given int) int {
	var value = 1
	for count := 1; count <= given; count++ {
		value *= count
	}
	return value
}
