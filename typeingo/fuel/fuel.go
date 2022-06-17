package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

//type Title string

// func ToGallons(l Liters) Gallons {
// 	return Gallons(l * 0.264)
// }

// func (l Liters) ToGallons() Gallons {
// 	return Gallons(l * 0.264)
// }

// func (m Milliliters) ToGallons() Gallons {
// 	return Gallons(m * 0.000264)
// }

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

// func ToLiters(g Gallons) Liters {
// 	return Liters(g * 3.785)
// }

func main() {
	//var carFuel Gallons
	//var busFuel Liters
	//carFuel = Gallons(Liters(40.0) * 0.264)
	//busFuel = Liters(Gallons(63.0) * 3.785)
	//fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel)

	// carFuel := Gallons(1.2)
	// busFuel := Liters(4.5)
	// carFuel += ToGallons(Liters(40.0))
	// busFuel += ToLiters(Gallons(30.0))
	// fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	// fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)

	// soda := Liters(2)
	// fmt.Printf("%0.3f liters quals %0.3f gallons\n", soda, soda.ToGallons())
	// water := Milliliters(500)
	// fmt.Printf("%0.3f millimetrs quals %0.3f glallons\n", water, water.ToGallons())

	milk := Gallons(2)
	fmt.Printf("%0.3f gallos quals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallos quals %0.3f milliliters\n", milk, milk.ToMilliliters())

}
