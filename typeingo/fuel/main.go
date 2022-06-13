package main

import "fmt"

type Liters float64
type Gallons float64

//type Title string

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}
func main() {
	//var carFuel Gallons
	//var busFuel Liters
	//carFuel = Gallons(Liters(40.0) * 0.264)
	//busFuel = Liters(Gallons(63.0) * 3.785)
	//fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)

}
