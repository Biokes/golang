package main

import "fmt"

func main() {
	number, _ := fmt.Println("enter  number: ")
	number1, _ := fmt.Println("enter  number: ")
	if number > number1 {
		fmt.Println("0")
	}
	if number < number1 {
		fmt.Println("1")
	}
	if number == number1 {
		fmt.Println("-1")
	}
}
