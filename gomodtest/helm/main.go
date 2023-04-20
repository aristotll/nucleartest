package main

import (
	"fmt"

	"helm.sh/helm/v3/pkg/chart/loader"
)

func main() {
	chart, err := loader.Load("mychart")
	if err != nil {
		panic(err)
	}
	chart.
	chart.Templates
	fmt.Printf("chart.Values: %v\n", chart.Values)
	//fmt.Printf("%+v\n", chart)
}
