package main

import (
	"../../Hologram-Go"
	"fmt"
	"strconv"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the Product Categories example")

	var product = HologramGo.GetProductCategories()

	fmt.Print("Product ID: ")
	fmt.Println(strconv.Itoa(product.GetProductId()))

	fmt.Println("Product Name: " + product.GetProductName())
}
