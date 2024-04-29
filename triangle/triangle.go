package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter your number: ")
	_, _ = fmt.Scanln(&number)
	for count := 0; count < number; count++ {
		for counter := 0; counter <= count; counter++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
