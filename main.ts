/**
 * Shopping Cart Program
 * Author: Joshua Adeyemi
 * Version: 1.0.0
 * Date: 2025-12-23
 */

let numItems: number = 10; // Example: 10 items
const itemNames: string[] = new Array(numItems);
const itemPrices: number[] = new Array(numItems);

// Example inputs for demonstration (replace with actual prompts if using an input system)
itemNames[0] = "diapers"; itemPrices[0] = 50;
itemNames[1] = "milk"; itemPrices[1] = 30;
itemNames[2] = "cheese"; itemPrices[2] = 25;
itemNames[3] = "steak"; itemPrices[3] = 60;
itemNames[4] = "rice"; itemPrices[4] = 40;
itemNames[5] = "ghee"; itemPrices[5] = 35;
itemNames[6] = "green peppers"; itemPrices[6] = 15;
itemNames[7] = "salmon"; itemPrices[7] = 55;
itemNames[8] = "pasta"; itemPrices[8] = 18;
itemNames[9] = "onions"; itemPrices[9] = 10;

// Calculate subtotal
let subtotal: number = 0;
for (let i = 0; i < numItems; i++) {
  subtotal += itemPrices[i];
}

// Check for discount
let discount: number = 0;
if (subtotal > 350) {
  discount = subtotal * 0.1;
}

// Subtotal after discount
let subtotalAfterDiscount: number = subtotal - discount;

// Calculate HST
let hst: number = subtotalAfterDiscount * 0.13;

// Calculate total
let total: number = subtotalAfterDiscount + hst;

// Output
console.log(`Your shopping cart includes: ${itemNames.join(", ")}`);
console.log(`The total amount of items in your shopping cart is ${numItems}\n`);
console.log(`The subtotal cost of your shopping trip was $${subtotal.toFixed(2)}`);
if (discount > 0) {
  console.log(`You are eligible for a discount of $${discount.toFixed(2)}`);
}
console.log(`The HST is $${hst.toFixed(2)}`);
console.log(`The total is $${total.toFixed(2)}`);

console.log("\nDone.");
