package main

import (
	"fmt"
)

func main() {
	fmt.Println(collectInput())
}

func collectInput() float64 {
	number := 0
	for number < 1 {
		fmt.Println("Enter a number to get its maths value!!: ")
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if number < 1 {
			fmt.Println("Please enter a number greater than or equal to 1.")
		}
	}
	return getConstant(number)
}

func getConstant(number int) float64 {
	total := 0.0
	for i := 0; i <= number; i++ {
		total += 1.0 / factorial(float64(i))
	}
	return total
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	fact := 1.0
	for i := 1; i <= int(n); i++ {
		fact *= float64(i)
	}
	return fact
}
