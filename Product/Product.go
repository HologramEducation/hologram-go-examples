package main

import (
	"../../Hologram-Go"
	"fmt"
)

// This example shows how to get all a product by id in our store.

func main() {

	// Initializing username and password.
	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the GetProduct example")

	// Query for product by productid
	var product = HologramGo.GetProduct(1)

	fmt.Println("Product SKU: " + product.GetProductSku())

	fmt.Println("Product Name: " + product.GetProductName())

	fmt.Println("Product Description: " + product.GetProductDescription())

	fmt.Print("Product Price: ")
	fmt.Println(product.GetProductPrice())

	fmt.Println("Product Image URL: " + product.GetProductImageUrl())

	fmt.Println("Product Invoice Description: " + product.GetProductInvoiceDescription())
}

// Example output:

// Welcome to the GetProduct example
// Product SKU: SIM
// Product Name: SIM Card
// Product Description: The Hologram Global SIM Card. Comes in a variety of sizes for any application. (Pick your data plan when you go to activate the SIM.)
// Product Price: 5.00
// Product Image URL: https://dashboard.hologram.io/img/sim_card.jpg
// Product Invoice Description: Global SIM Card
