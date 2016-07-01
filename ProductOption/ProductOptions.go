package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the Product Options example")

	var product = HologramGo.Product{}

	product.GetProductOptions()
}
