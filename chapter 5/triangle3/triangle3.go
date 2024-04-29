package main

import "fmt"

func main() {
	for count := 0; count < 10; count++ {
		for counter := 0; counter < count; counter++ {
			fmt.Print(" ")
		}
		for index := 0; index < 10-count; index++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
