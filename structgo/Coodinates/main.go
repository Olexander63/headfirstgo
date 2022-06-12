package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/structgo/geo"
)

func main() {
	location := geo.Landmark{}
	location.Name = " The Googolplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
