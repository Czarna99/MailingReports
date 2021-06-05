package main

import (
	app "MailingReports/MailingReports/App"
	"fmt"
)

var i string

func main() {
	fmt.Println("Welcome in automatic reports sending program.\n\n Please pick your choice:")
	fmt.Println("t - send reports, e - exit")
	fmt.Scanf("%s", &i)

	switch i {

	case "t":

		app.ApplicationStart()
		return
	case "e":
		fmt.Println("Program closed. Press enter to continue.")
		return
	default:
		fmt.Println("Incorrect choice. Program will exit.")
	}

}
