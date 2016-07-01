package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the Product Categories example")

	var product = HologramGo.GetProductCategories()

	productCategory := (HologramGo.ProductCategory)(product[0].(map[string]interface{}))

	fmt.Print("Product Category Name: ")
	fmt.Println(productCategory.GetProductCategoryName())
}

// Example output:

// Welcome to the Product Categories example
// Product Category Name: accessories
