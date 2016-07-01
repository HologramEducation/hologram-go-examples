package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the GetDevice example")

	// Get a device with device id 37739
	var device = HologramGo.GetDevice(37739)

	fmt.Print("Device ID: ")
	fmt.Println(device.GetDeviceId())

	fmt.Print("User ID: ")
	fmt.Println(device.GetDeviceUserId())

	fmt.Println("Device Name: " + device.GetDeviceName())

	fmt.Println("Device Type: " + device.GetDeviceType())

	fmt.Println("Device Created: " + device.GetWhenCreated())
}

// Example output:

// Welcome to the GetDevice example
// Device ID: 37739
// User ID: 2642
// Device Name: Unnamed Device (07873)
// Device Type: Unknown
// Device Created: 2016-06-06 01:19:02
