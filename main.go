package main

import "fmt"

func main() {
	// Items in the cart
	items := [10]string{"diapers", "milk", "cheese", "steak", "rice", "ghee", "green peppers", "salmon", "pasta", "onions"}
	prices := [10]float64{45.00, 35.00, 50.00, 60.00, 30.00, 28.00, 25.00, 40.00, 25.45, 60.00}

	// Calculate subtotal
	subtotal := 0.0
	for i := 0; i < 10; i++ {
		subtotal += prices[i]
	}

	// Calculate discount
	discount := 0.0
	if subtotal > 350.0 {
		discount = subtotal * 0.10
	}

	// Calculate HST (13% of subtotal after discount)
	hst := (subtotal - discount) * 0.13

	// Calculate total
	total := subtotal - discount + hst

	// Output
	fmt.Print("Your shopping cart includes: ")
	for i := 0; i < 10; i++ {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(items[i])
	}
	fmt.Println()
	fmt.Printf("The total amount of items in your shopping cart is %d\n\n", len(items))
	fmt.Printf("The subtotal cost of your shopping trip was $%.2f\n", subtotal)
	if discount > 0 {
		fmt.Printf("You are eligible for a discount of $%.2f\n", discount)
	}
	fmt.Printf("The HST is $%.2f\n", hst)
	fmt.Printf("The total is $%.2f\n\n", total)
	
	fmt.Println("Done.")
}
