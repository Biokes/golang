package main

import "fmt"

func main() {
	var number int
	total := 0
	value := 0
	fmt.Println("Enter a number: ")
	number, _ = fmt.Scanln(&number)
	for total <= number {
		fmt.Println("Enter a number: ")
		number, _ = fmt.Scanln(&number)
		value += number
	}
	fmt.Println(number)
}