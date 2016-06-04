package main

import (
	"../Hologram-Go"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the GetDevice example")

	var device = HologramGo.Device{}
	
	device.GetDevice(10)
}
