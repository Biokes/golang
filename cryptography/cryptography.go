package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func encryptNumber(numberGiven string) string {
	var num = (int(numberGiven[0]-'0') + 7) % 10
	var num2 = (int(numberGiven[1]-'0') + 7) % 10
	var num3 = (int(numberGiven[2]-'0') + 7) % 10
	var num4 = (int(numberGiven[3]-'0') + 7) % 10
	var output = fmt.Sprintf("%d%d%d%d", num3, num4, num, num2)
	return output
}
func validate(value int) string {
	var number int
	character := fmt.Sprintf("%d", value)
	for len(character) != 4 {
		fmt.Println("Enter a four-digit number!!! :")
		number, _ = fmt.Scanln(&number)
	}
	return fmt.Sprintf("%d", number)
}
func getInput() string {
	var input int
	fmt.Println("enter a 4digit number to encrypt: ")
	input, _ = fmt.Scanln(&input)
	return validate(input)
}
func decrypt(value string) string {
	validateString(value)
	num1 := parse(strconv.Itoa(int(value[0])))
	num2 := parse(strconv.Itoa(int(value[1])))
	num3 := parse(strconv.Itoa(int(value[2])))
	num4 := parse(strconv.Itoa(int(value[3])))
	return fmt.Sprintf("%d%d%d%d", num4, num3, num1, num2)
}
func parse(value string) any {
	number, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Sprintf("an %s occured due to %s given\npls check and retry", "error", value)
	}
	return number
}
func validateString(given string) string {
	var value int
	for len(given) != 4 || (reflect.TypeOf(given) != reflect.TypeOf("")) {
		fmt.Println("Enter a four-digit number for decryption: ")
		value, _ = fmt.Scanln(&value)
	}
	if len(given) != 4 || (reflect.TypeOf(given) != reflect.TypeOf("")) {
		return fmt.Sprintf("%d", value)
	}
	return given
}
func main() {
	number := getInput()
	str := encryptNumber(number)
	value := decrypt(str)
	fmt.Println(fmt.Sprintf("Encrypt: %s\nDecrypt : %s ", number, value))
}
