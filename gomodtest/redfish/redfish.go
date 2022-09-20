package main

import (
	"context"
	"fmt"
	redfish "github.com/Nordix/go-redfish/client"
)

func main() {
	cfg := &redfish.Configuration{
		BasePath:      "10.0.120.13",
		DefaultHeader: make(map[string]string),
		UserAgent:     "go-redfish/client",
	}

	redfishApi := redfish.NewAPIClient(cfg).DefaultApi
	sl, _, _ := redfishApi.ListSystems(context.Background())
	fmt.Println(sl)
}
