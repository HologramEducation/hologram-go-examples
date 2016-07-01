package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the Data plan by Plan ID example")

	// Get a data plan with device id 37739
	var dataPlan = HologramGo.GetDeviceDataPlan(51)

	fmt.Print("Data Plan ID: ")
	fmt.Println(dataPlan.GetDataPlanId())

	fmt.Print("Data Plan Partner ID: ")
	fmt.Println(dataPlan.GetDataPlanPartnerId())

	fmt.Println("Data Plan Name: " + dataPlan.GetDataPlanName())

	fmt.Println("Data Plan Description: " + dataPlan.GetDataPlanDescription())

	fmt.Print("Data Plan Size: ")
	fmt.Println(dataPlan.GetDataPlanSize())
}

// Expected output:

// Welcome to the Data plan by Plan ID example
// Data Plan ID: 51
// Data Plan Partner ID: 1
// Data Plan Name: 1 MB
// Data Plan Description:
// Data Plan Size: 1e+06
