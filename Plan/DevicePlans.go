package main

import (
	"../../Hologram-Go"
	"fmt"
)

func main() {

	HologramGo.InitializeUsernameAndPassword("../credentials.json");

	fmt.Println("Welcome to the Device Plans example")

	// Get a device with device id 37739
	plans := HologramGo.GetDeviceDataPlans()

	for _, plan := range plans {
	    
	    dataPlan := (HologramGo.Plan)(plan.(map[string]interface{}))

	    fmt.Print("Data Plan ID: ")
	    fmt.Println(dataPlan.GetDataPlanId())

	    fmt.Print("Data Plan Partner ID: ")
	    fmt.Println(dataPlan.GetDataPlanPartnerId())

	    fmt.Println("Data Plan Name: " + dataPlan.GetDataPlanName())

	    fmt.Println("Data Plan Description: " + dataPlan.GetDataPlanDescription())

	    fmt.Print("Data Plan Size: ")
	    fmt.Println(dataPlan.GetDataPlanSize())
	}
}

// Example output:

// Welcome to the Device Plans example
// Data Plan ID: 101
// Data Plan Partner ID: 99
// Data Plan Name: TEMP PLAN
// Data Plan Description:
// Data Plan Size: 0
// Data Plan ID: 100
// Data Plan Partner ID: 99
// Data Plan Name: TEMP PLAN
// Data Plan Description:
// Data Plan Size: 0
// Data Plan ID: 80
// Data Plan Partner ID: 1
// Data Plan Name: PAYG
// Data Plan Description:
// Data Plan Size: 0
// Data Plan ID: 79
// Data Plan Partner ID: 1
// Data Plan Name: PAYG
// Data Plan Description:
// Data Plan Size: 0
// Data Plan ID: 78
// Data Plan Partner ID: 1
// Data Plan Name: 2 MB
// Data Plan Description:
// Data Plan Size: 2e+06
// Data Plan ID: 77
// Data Plan Partner ID: 1
// Data Plan Name: 1 MB + 30-day Free Trial
// Data Plan Description:
// Data Plan Size: 1e+06
// Data Plan ID: 75
// Data Plan Partner ID: 1
// Data Plan Name: 5 MB
// Data Plan Description:
// Data Plan Size: 5e+06
// Data Plan ID: 74
// Data Plan Partner ID: 1
// Data Plan Name: 5 MB
// Data Plan Description:
// Data Plan Size: 5e+06
// Data Plan ID: 73
// Data Plan Partner ID: 1
// Data Plan Name: 500 MB
// Data Plan Description:
// Data Plan Size: 5e+08
// Data Plan ID: 72
// Data Plan Partner ID: 1
// Data Plan Name: 250 MB
// Data Plan Description:
// Data Plan Size: 2.5e+08
// Data Plan ID: 71
// Data Plan Partner ID: 1
// Data Plan Name: 125 MB
// Data Plan Description:
// Data Plan Size: 1.25e+08
// Data Plan ID: 70
// Data Plan Partner ID: 1
// Data Plan Name: 100 MB
// Data Plan Description:
// Data Plan Size: 1e+08
// Data Plan ID: 69
// Data Plan Partner ID: 1
// Data Plan Name: 50 MB
// Data Plan Description:
// Data Plan Size: 5e+07
// Data Plan ID: 68
// Data Plan Partner ID: 1
// Data Plan Name: 30 MB
// Data Plan Description:
// Data Plan Size: 3e+07
// Data Plan ID: 67
// Data Plan Partner ID: 1
// Data Plan Name: 20 MB
// Data Plan Description:
// Data Plan Size: 2e+07
// Data Plan ID: 66
// Data Plan Partner ID: 1
// Data Plan Name: 10 MB
// Data Plan Description:
// Data Plan Size: 1e+07
// Data Plan ID: 65
// Data Plan Partner ID: 1
// Data Plan Name: 5 MB
// Data Plan Description:
// Data Plan Size: 5e+06
// Data Plan ID: 64
// Data Plan Partner ID: 1
// Data Plan Name: 3 MB
// Data Plan Description:
// Data Plan Size: 3e+06
// Data Plan ID: 62
// Data Plan Partner ID: 1
// Data Plan Name: 1 MB + 30-day Free Trial
// Data Plan Description:
// Data Plan Size: 1e+06
// Data Plan ID: 61
// Data Plan Partner ID: 1
// Data Plan Name: 500 KB
// Data Plan Description:
// Data Plan Size: 500000
// Data Plan ID: 57
// Data Plan Partner ID: 1
// Data Plan Name: 10 MB
// Data Plan Description:
// Data Plan Size: 1e+07
// Data Plan ID: 55
// Data Plan Partner ID: 1
// Data Plan Name: 5 MB
// Data Plan Description:
// Data Plan Size: 5e+06
// Data Plan ID: 52
// Data Plan Partner ID: 1
// Data Plan Name: 3 MB
// Data Plan Description:
// Data Plan Size: 3e+06
// Data Plan ID: 51
// Data Plan Partner ID: 1
// Data Plan Name: 1 MB
// Data Plan Description:
// Data Plan Size: 1e+06
// Data Plan ID: 49
// Data Plan Partner ID: 1
// Data Plan Name: 10 MB
// Data Plan Description:
// Data Plan Size: 1e+07