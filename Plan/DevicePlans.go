package main

import (
	"../../Hologram-Go"
	"fmt"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json");

	fmt.Println("Welcome to the Device Plans example")

	var plan = HologramGo.Plans{}

	plan.GetDeviceDataPlans()
}
