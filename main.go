package main

import (
	"fmt"
)

func main() {
	itemNames := [10]string{"diapers", "milk", "cheese", "steak", "rice", "ghee", "green peppers", "salmon", "pasta", "onions"}
	itemPrices := [10]float64{30.5, 25.0, 40.0, 55.0, 18.0, 20.0, 25.0, 45.0, 40.0, 0.0}

	numItems := len(itemNames)

	subtotal := 0.0
	for i := 0; i < numItems; i++ {
		subtotal += itemPrices[i]
	}

	discount := 0.0
	if subtotal > 350 {
		discount = subtotal * 0.1
	}

	subtotalAfterDiscount := subtotal - discount
	hst := subtotalAfterDiscount * 0.13
	total := subtotalAfterDiscount + hst

	fmt.Printf("Your shopping cart includes: ")
	for i, name := range itemNames {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(name)
	}
	fmt.Printf("\nThe total amount of items in your shopping cart is %d\n\n", numItems)
	fmt.Printf("The subtotal cost of your shopping trip was $%.2f\n", subtotal)
	if discount > 0 {
		fmt.Printf("You are eligible for a discount of $%.2f\n", discount)
	}
	fmt.Printf("The HST is $%.2f\n", hst)
	fmt.Printf("The total is $%.2f\n", total)
	fmt.Println("\nDone.")
}
