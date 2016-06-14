package main

import (
	"../../Hologram-Go"
	"fmt"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the Product Options example")

	var product = HologramGo.Product{}

	product.GetProductOptions()
}
