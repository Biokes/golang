package main

import "fmt"

func getNumbers(times int) int {
	min := 2147483647
	max := -2147483647
	var num int
	for counter := 1; counter <= times; counter++ {
		fmt.Println("Enter number: ")
		_, _ = fmt.Scanln(&num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min + max
}
func main() {
	var num int
	fmt.Println("Enter number: ")
	_, _ = fmt.Scanln(&num)
	fmt.Println(getNumbers(num))
}
