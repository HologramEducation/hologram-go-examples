package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

// This example will hopefully give you an idea on how to query your API key via
// our REST API.

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the API Key example")

	var user = HologramGo.GetAPIKey()

	fmt.Println("User API Key: " + user.GetUserAPIKey())

	// TODO: API Key generation.
}
