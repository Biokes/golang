package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var num int
	_, _ = fmt.Scanln(&num)
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	fmt.Println(total)
}
