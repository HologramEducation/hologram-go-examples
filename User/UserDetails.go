package main

import (
	"fmt"
	"github.com/hologram-io/hologram-go"
)

// This example will hopefully give you an idea on how to query user details via
// our REST API.

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json")

	fmt.Println("Welcome to the User Details example")

	var user = HologramGo.GetAPIKey()

	fmt.Println("User API Key: " + user.GetUserAPIKey())
}
