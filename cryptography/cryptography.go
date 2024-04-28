package main

import (
	"fmt"
	"strconv"
)

func wrapToInteger(value string) int {
	output, _ := strconv.ParseInt(value, 10, 32)
	return int(output)
}
func encryptNumber(numberGiven string) string {
	index1 := (wrapToInteger(string(numberGiven[0])) + 7) % 10
	index2 := (wrapToInteger(string(numberGiven[1])) + 7) % 10
	index3 := (wrapToInteger(string(numberGiven[2])) + 7) % 10
	index4 := (wrapToInteger(string(numberGiven[3])) + 7) % 10
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
func decrypt(input string) int {
	return (wrapToInteger(input) + 3) % 10
}
func decryptGiven(value string) string {
	index1 := decrypt(string(value[0]))
	index2 := decrypt(string(value[1]))
	index3 := decrypt(string(value[2]))
	index4 := decrypt(string(value[3]))
	return fmt.Sprintf("%d%d%d%d", index3, index4, index1, index2)
}
func encrypt() string {
	input := getInput()
	validate(input)
	return encryptNumber(input)
}
func main() {
	input := encrypt()
	value := decryptGiven(input)
	fmt.Println(fmt.Sprintf("Encrypt: %s\nDecrypt : %s ", input, value))
}
