package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	var num1 int
	var num2 int
	var num3 int
	fmt.Println("Enter a set of points: : ")
	number, _ = fmt.Scanln(&number)
	num1, _ = fmt.Scanln(&num1)
	num2, _ = fmt.Scanln(&num2)
	num3, _ = fmt.Scanln(&num3)
	fmt.Println(fmt.Println(calculatePoints(number, num1, num2, num3)))
}
func calculatePoints(num1 int, num2 int, num3 int, num4 int) float32 {
	var point1 = (num2 - num1) * (num2 - num1)
	var point2 = (num4 - num3) * (num4 - num3)
	return float32(math.Pow(float64(point2+point1), float64(1/2)))
}
