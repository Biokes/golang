package main

import "fmt"

func main() {
	var number int
	total := 0
	value := 0
	fmt.Println("Enter a number: ")
	number, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)
	}
	value = number
	for total < value {
		fmt.Println("Enter a number: ")
		_, _ = fmt.Scanln(&number)
		total += number
	}
	fmt.Println(number)
}
