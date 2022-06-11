package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/structgo/magazine"
)

//type subscriber struct {
//	name   string
//	rate   float64
//	active bool
//}

//func printInfo(s *subscriber) {
//	fmt.Println("Name:", s.name)
//	fmt.Println("Monthly rate:", s.rate)
//	fmt.Println("Active?", s.active)
//}

//func defaultSubscriber(name string) *subscriber {
//	var s subscriber
//	s.name = name
//	s.rate = 5.99
//	s.active = true
//	return &s
//}

//func applyDiscounts(s *subscriber) {
//	s.rate = 4.99
//}

//type myStruct struct {
//	myField int
//}

func main() {
	//subscriber1 := defaultSubscriber("Aman Singh")
	//applyDiscounts(subscriber1)
	//printInfo(subscriber1)
	//
	//subscriber2 := defaultSubscriber("Seth Ryan")
	//printInfo(subscriber2)

	//subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	//fmt.Println("Name", subscriber.Name)
	//fmt.Println("Rate", subscriber.Rate)
	//fmt.Println("Active: ", subscriber.Active)

	//var employee magazine.Employee
	//employee.Name = "Joy Carr"
	//employee.Salary = 60000
	//fmt.Println(employee.Name)
	//fmt.Println(employee.Salary)

	//var address magazine.Address
	//address.Street = "123 Oak St"
	//address.City = "Omaha"
	//address.State = "Ne"
	//address.PostalCode = "68111"
	//fmt.Println(address)
	//address := magazine.Address{Street: "123 Oa kSt",
	//	City: "Omaha", State: "Ne", PostalCode: "68111"}
	//subscriber := magazine.Subscriber{Name: "Aman Singh"}
	//subscriber.HomeAddress = address

	//subscriber := magazine.Subscriber{}
	//fmt.Printf("%#v\n", subscriber.HomeAddress)
	//
	//subscriber.HomeAddress.PostalCode = "68111"
	//fmt.Println("postal Code:", subscriber.HomeAddress.PostalCode)

	//var s subscriber
	//applyDiscounts(&s)
	//fmt.Println(s.rate)

	//var value int = 2
	//var pointer *int = &value
	//fmt.Println(*pointer)

	//var value myStruct
	//value.myField = 3
	//var pointer *myStruct = &value
	//pointer.myField = 9
	//fmt.Println(pointer.myField)

	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"

	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Address.Street)
	fmt.Println("City:", subscriber.Address.City)
	fmt.Println("State:", subscriber.Address.State)
	fmt.Println("Postal Code", subscriber.Address.PostalCode)

	employee := magazine.Employee{Name: "Joy Car"}
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Postal Code:", employee.Address.PostalCode)
}
