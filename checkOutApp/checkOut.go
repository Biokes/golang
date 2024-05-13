package main

import (
	"fmt"
	"strings"
	"time"
)

func getItemBought() string {
	var item string
	fmt.Println("What did the customer buy: ")
	_, _ = fmt.Scanln(&item)
	for len(strings.TrimSpace(item)) == 0 {
		fmt.Println("Enter a valid item name!!!\nWhat did the customer buy: ")
		_, _ = fmt.Scanln(&item)
	}
	return item
}
func getNumberOfUnits() int {
	var item int
	fmt.Println("How many pieces: ")
	_, _ = fmt.Scanln(&item)
	for item <= 0 {
		fmt.Println("Enter a valid amount of item Gotten!!!\nHow many pieces: ")
		_, _ = fmt.Scanln(&item)
	}
	return item
}
func getPrice() float32 {
	var item float32
	fmt.Println("How much per unit: ")
	_, _ = fmt.Scanln(&item)
	for item <= 0 {
		fmt.Println("Enter a valid amount for price per unit!!!\nHow much per unit: ")
		_, _ = fmt.Scanln(&item)
	}
	return item
}
func getName() string {
	fmt.Println("Enter customer's name: ")
	var name string
	_, _ = fmt.Scanln(&name)
	return name
}
func getDetails() ([]string, []int, []float32) {
	var items []string
	var pieces []int
	var prices []float32
	input := "no"
	for !strings.EqualFold(input, "yes") {
		items = append(items, getItemBought())
		pieces = append(pieces, getNumberOfUnits())
		prices = append(prices, getPrice())
		fmt.Println("Add more item: ")
		_, _ = fmt.Scanln(&input)
	}
	return items, pieces, prices
}
func getNameAndDiscount(customerName string) (string, float32) {
	var name string
	var discount float32
	fmt.Println("Enter your name: ")
	_, _ = fmt.Scanln(&name)
	fmt.Println(`How much discount will ` + customerName + `get: `)
	_, _ = fmt.Scanln(&discount)
	return name, discount
}
func bill(customerName string, cashierName string, discount float32, goods []string, units []int, pricePerUnit []float32) float32 {
	fmt.Println("Semicolon Store")
	fmt.Sprintf(`Main Branch\nLocation : My House\nTel : 012901928301\nDate: %s`, time.Now().Format("02/01/2006 15:04 PM"))
	fmt.Sprintf(`cashier's name: %s\nCustomer's name: %s`, customerName, cashierName)
	printLines()
	fmt.Sprintf("%10s %5s %5s %9s", "ITEM", "QTY", "PRICE", "Total(NGN)")
	printBrokenLines()
	var total []float32
	for count := 0; count < len(goods); count++ {
		total = append(total, float32(units[count])*pricePerUnit[count])
		fmt.Sprintf("%10s %5d %5.2f %9.2f\n", goods[count], units[count], pricePerUnit[count], total[count])
	}
	var subTotal float32
	for count := 0; count < len(total); count++ {
		subTotal += total[count]
	}
	printBrokenLines()
	fmt.Sprintf("%20s : %10.2f", "Sub total", subTotal)
	fmt.Sprintf("%20s : %10.2f", "Discount", discount)
	fmt.Sprintf("%20s : %10.2f", "VAT@ 17.5%", subTotal*0.175)
	printLines()
	fmt.Sprintf("%20s : %10.2f", "Bill Total", subTotal-(subTotal*discount)/100+(subTotal*0.175))
	printLines()
	fmt.Sprintf("this is not a reciept. Kindly pay %.2f", subTotal-(subTotal*discount)/100+(subTotal*0.175))
	return subTotal - (subTotal*discount)/100 + (subTotal * 0.175)
}
func printLines() string {
	return "========================================================================"
}
func printBrokenLines() string {
	return "-------------------------------------------------------------------------"

}
func payment(amount float32, customerName string, cashierName string, discount float32, goods []string, units []int, pricePerUnit []float32) {
	fmt.Println("Semicolon Store")
	fmt.Sprintf(`Main Branch\nLocation : My House\nTel : 012901928301\nDate: %s`, time.Now().Format("02/01/2006 15:04 PM"))
	fmt.Sprintf(`cashier's name: %s\nCustomer's name: %s`, customerName, cashierName)
	printLines()
	fmt.Sprintf("%10s %5s %5s %9s", "ITEM", "QTY", "PRICE", "Total(NGN)")
	printBrokenLines()
	var total []float32
	for count := 0; count < len(goods); count++ {
		total = append(total, float32(units[count])*pricePerUnit[count])
		fmt.Sprintf("%10s %5d %5.2f %9.2f\n", goods[count], units[count], pricePerUnit[count], total[count])
	}
	var subTotal float32
	for count := 0; count < len(total); count++ {
		subTotal += total[count]
	}
	printBrokenLines()
	fmt.Sprintf("%20s : %10.2f", "Sub total", subTotal)
	fmt.Sprintf("%20s : %10.2f", "Discount", discount)
	fmt.Sprintf("%20s : %10.2f", "VAT@ 17.5%", subTotal*0.175)
	printLines()
	fmt.Sprintf("%20s : %10.2f", "Bill total", subTotal)
	fmt.Sprintf("%20s : %10.2f", "Amount paid", amount)
	fmt.Sprintf("%20s : %10.2f", "Balance")
	printLines()
	fmt.Println("Thanks for your patronage.")
	printLines()
}
func getAmount(total float32) float32 {
	var amount float32
	fmt.Println("How much did the customer pay: ")
	_, _ = fmt.Scanln(amount)
	for amount < total {
		fmt.Sprintf("Too small ask for at least %.2f\nHow much did the customer pay: ", amount)
		_, _ = fmt.Scanln(amount)
	}
	return amount
}
func checkOut() {
	name := getName()
	items, units, prices := getDetails()
	cashier, discount := getNameAndDiscount(name)
	total := bill(name, cashier, discount, items, units, prices)
	fmt.Sprintln("\n\n")
	amountPaid := getAmount(total)
	fmt.Sprintln("\n\n")
	payment(amountPaid, name, cashier, discount, items, units, prices)
}
func main() {
	checkOut()
}
