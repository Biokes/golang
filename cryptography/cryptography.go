package main

import (
	"fmt"
	"strconv"
)

func encryptNumber(numberGiven string) string {
	index1, _ := strconv.ParseInt(numberGiven[0], 10, 64)
	index2 := numberGiven[1]
	index3 := numberGiven[2]
	index4 := numberGiven[3]
	return fmt.Sprintf("%d%d%d%d", index3, index4, index1, index2)
}
func validate(value string) string {
	for len(value) != 4 {
		fmt.Println("enter a 4-digit number to encrypt: ")
		_, _ = fmt.Scanln(&value)
	}
	return value
}
func getInput() string {
	var input string
	fmt.Println("enter a 4-digit number to encrypt: ")
	_, _ = fmt.Scanln(&input)
	return validate(input)
}
func decrypt(value string) string {

	return ""
}
func main() {
	number := getInput()
	str := encryptNumber(number)
	value := decrypt(str)
	fmt.Println(fmt.Sprintf("Encrypt: %s\nDecrypt : %s ", number, value))
}
