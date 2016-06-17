package main

import (
	"../../Hologram-Go"
	"fmt"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to this example where you can claim ownership and activate a device")

	// Claim and activate a device with sim id 
	var device = HologramGo.ClaimOwnershipAndActivateDevice(8944501410155200039)

	fmt.Print("Device ID: ")
	fmt.Println(device.GetDeviceId())

	fmt.Print("User ID: ")
	fmt.Println(device.GetDeviceUserId())

	fmt.Println("Device Name: " + device.GetDeviceName())

	fmt.Println("Device Type: " + device.GetDeviceType())

	fmt.Println("Device Created: " + device.GetWhenCreated())
}

// Expected output:
