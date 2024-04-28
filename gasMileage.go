package main

import "fmt"

func getGallonsUsed() int {
	var gallons int
	fmt.Println("Enter number of gallons used: ")
	var _, _ = fmt.Scanln(&gallons)
	return gallons
}
func displayGallonsUsed() float32 {
	var number int
	fmt.Println("Enter total number of miles: ")
	_, _ = fmt.Scanln(&number)
	gallonsUsed := getGallonsUsed()
	return float32(number) / float32(gallonsUsed)
}
func display() {
	choice := 1
	for choice != -1 {
		var total = displayGallonsUsed()
		fmt.Println("total miles/Gallon => ", total, "mph")
		fmt.Println("Enter any number to continue and -1 to stop: ")
		var _, _ = fmt.Scanln(&choice)
	}
}
func main() {
	display()
}
