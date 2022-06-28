package main

import (
	"fmt"
	calendar "github.com/Olexander63/headfistgo/encapsulation/date"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2022)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
