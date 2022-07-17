package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/interfacesgo/testinterface/mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
