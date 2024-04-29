package main

import "fmt"

func collectInput() string {
	var input string
	fmt.Print("Enter a five digit number: ")
	_, _ = fmt.Scanln(&input)
	for len(input) != 5 {
		fmt.Print("Enter a five digit number: ")
		_, _ = fmt.Scanln(&input)
	}
	return input
}
func isPalindrome(value string) bool {
	palindrome := fmt.Sprintf("%s%s%s%s%s", string(value[4]), string(value[3]),
		string(value[2]), string(value[1]), string(value[0]))
	return value == palindrome
}
func main() {
	input := collectInput()
	fmt.Printf(input+"%s is a palindrome: ", isPalindrome(input))
}
