package main

import "fmt"

func main() {
	var number int
	total := 0
	value := 0
	fmt.Println("Enter a number: ")
	_, _ = fmt.Scanln(&number)
	for total <= value {
		fmt.Println("Enter a number: ")
		number, _ = fmt.Scanln(&number)
		value += number
	}
	fmt.Println(number)
}
