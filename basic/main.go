package main

import "fmt"

func main() {
	////hello world
	//fmt.Println("Hello")
	//fmt.Println(math.Floor(2.75))
	//fmt.Println(strings.Title("head first go"))
	//
	////reflect
	//
	//fmt.Println(reflect.TypeOf(42))
	//fmt.Println(reflect.TypeOf(3.1415))
	//fmt.Println(reflect.TypeOf(true))
	//fmt.Println(reflect.TypeOf("Hello, GO"))

	//customer
	//quantity := 4
	//length, width := 1.2, 2.4
	//customerName := "Damon Cole"

	//fmt.Println(customerName)
	//fmt.Println("has ordered", quantity, "sheets")
	//fmt.Println("each with an area of")
	//fmt.Println(length*width, "square meters")

	//casting
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))

	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))
	//taxRate end

}
