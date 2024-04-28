package main

import "fmt"

func main() {
	year := 2022
	population := 7951000000
	for i := 1; i <= 75; i++ {
		fmt.Println("Year population")
		fmt.Printf("%d %-2d\n", year, population)
		year += 1
		population += calculatePopulation(population)
	}
}
func calculatePopulation(population int) int {
	return (31 * population) / 1000
}
