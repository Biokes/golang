package main

import "fmt"

func main() {
	year := 2022
	population := 7951000000
	populationIncrease := population
	for i := 1; i <= 75; i++ {
		fmt.Printf("%s %-3s   %s", "Year", "increase", "current increase")
		fmt.Printf("\n%d %3d %17d\n", year, population-populationIncrease, population)
		populationIncrease = population
		year += 1
		population += calculatePopulation(population)

	}
}
func calculatePopulation(population int) int {
	return (31 * population) / 1000
}
