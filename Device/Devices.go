package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the GetDevices example")

	var devices = HologramGo.GetDevices()

	device := (HologramGo.Device)(devices[0].(map[string]interface{}))

	fmt.Print("Device ID: ")
	fmt.Println(device.GetDeviceId())

	fmt.Print("User ID: ")
	fmt.Println(device.GetDeviceUserId())

	fmt.Println("Device Name: " + device.GetDeviceName())

	fmt.Println("Device Type: " + device.GetDeviceType())

	fmt.Println("Device Created: " + device.GetWhenCreated())
}

// Example output:

// Welcome to the GetDevices example
// Device ID: 84
// User ID: 4477
// Device Name: Unnamed Device (90123)
// Device Type: Unknown
// Device Created: 2016-06-06 21:51:16
