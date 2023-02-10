package main

import (
	"fmt"
	route2 "github.com/rodolfosouzamenezes/codelivery-simulator/app/route"
)

func main() {
	route := route2.Route{
		ID: "1",
		ClientID: "1",
	}
	
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[1])
}