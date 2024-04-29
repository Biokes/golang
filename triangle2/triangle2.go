package main

import "fmt"

func main() {
	for count := 0; count < 10; count++ {
		for index := 0; index < count; index++ {
			fmt.Println(" ")
		}
		for couter := 0; couter < 10-count; couter++ {
			fmt.Println("#")
		}
		fmt.Println()
	}
}
