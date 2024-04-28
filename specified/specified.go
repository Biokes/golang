package main

import "fmt"

func main() {
	var number int
	value := 0
	fmt.Println("Enter a number: ")
	number, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)
	}
	for number < value {
		fmt.Println("Enter a number: ")
		number, err = fmt.Scanln(&number)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(number)
}
