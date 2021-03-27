package main

import (
	"fmt"

	"github.com/MeztliRA/weekdays"
)

func main() {
	weekdays.PrintMessage()
	if weekdays.IsWeekday() {
		fmt.Println("Go to work!")
	} else {
		fmt.Println("Dont go to work")
	}
}
