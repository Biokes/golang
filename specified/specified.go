package main

import "fmt"

func main() {
	var number int
	value := 0
	fmt.Println("Enter a number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(err)
	}
	for value < number {
		var numb int
		fmt.Println("Enter a number: ")
		number, err = fmt.Scanln(&numb)
		if err != nil {
			fmt.Println(err)
		}
		value += numb
	}
	fmt.Printf("output : %d", number)
}
