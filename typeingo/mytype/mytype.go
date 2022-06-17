package main

import "fmt"

type MyType string

// func (m MyType) sayHi() {
// 	fmt.Println("Hi from", m)
// }

// func (m MyType) MethdoWhithParameters(number int, flag bool) {
// 	fmt.Println(m)
// 	fmt.Println(number)
// 	fmt.Println(flag)
// }

// func (m MyType) WithReturn() int {
// 	return len(m)
// }

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}
func main() {
	// value := MyType("a MyType value")
	// value.sayHi()
	// anotherValue := MyType("another value")
	// anotherValue.sayHi()
	// value := MyType("MyType value")
	// value.MethdoWhithParameters(4, true)

	value := MyType("a vaue")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
}
