package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/packagesexercise/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
