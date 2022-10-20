package main

import (
	"fmt"
	"mk"

	"github.com/TwiN/go-color"
)

func main() {
	mk.Cls()
	println(color.InRed(color.InBold("Welcome To My Program")))
	println(color.InBlue("----------------------------------"))
	var flag bool = true
	var choice int
	for flag {

		ShowMenu()
		println(color.InCyan("Select a Choice:"))
		fmt.Scan(&choice)

		switch choice {
		case 1:
			EnterData()
		case 2:
			Search()
		case 3:
			ShowData()
		case 4:
			flag = false
		default:
			mk.Cls()
			println(color.InPurple("Invalid Choice. Try Again."))
		}
	}

	println(color.InGreen("See You Later, Good Bye."))

}

var name []string
var bdate []string
var id_code []string

func EnterData() {
	mk.Cls()
	println(color.InBlue("Enter Your Name-Family:"))
	var str string
	fmt.Scan(&str)
	name = append(name, str)
	println(color.InBlue("Enter Your Birth Date:"))
	fmt.Scan(&str)
	bdate = append(bdate, str)

	println(color.InBlue("Enter Your ID Code:"))
	fmt.Scan(&str)
	id_code = append(id_code, str)

	mk.Cls()
	println(color.InGreen("Saved Successfully."))
	println(color.InYellow("----------------------------"))

}
func ShowMenu() {
	print(color.InGreen("1) "))
	println(color.InYellow("Data Entery"))

	print(color.InGreen("2) "))
	println(color.InYellow("Search"))

	print(color.InGreen("3) "))
	println(color.InYellow("Show Data"))

	print(color.InGreen("4) "))
	println(color.InYellow("Exit"))

}

func ShowData() {
	mk.Cls()
	for k, v := range name {
		fmt.Printf("Name[%d]=%s, ", k, v)
	}
	fmt.Println("")

	for k, v := range bdate {
		fmt.Printf("Birth Date[%d] = %s, ", k, v)
	}
	fmt.Println("")

	for k, v := range id_code {
		fmt.Printf("ID Code[%d] = %s, ", k, v)
	}
	fmt.Println("\n------------------------------")

	print(color.InBlue("Names =\t\t"))
	fmt.Println(name)

	print(color.InBlue("Birth Dates =\t"))
	fmt.Println(bdate)

	print(color.InBlue("IDs =\t\t"))
	fmt.Println(id_code)
	fmt.Println("------------------------------")

}

func Search() {
	mk.Cls()
	print(color.InGreen("1) "))
	println(color.InYellow("Search By Name"))

	print(color.InGreen("2) "))
	println(color.InYellow("Search By Birth Date"))

	print(color.InGreen("3) "))
	println(color.InYellow("Search By ID Code"))

	var choice int
	println(color.InCyan("Select a Choice:"))
	fmt.Scan(&choice)

	var found_item int = -1
	var selected_item string
	switch choice {
	case 1:
		mk.Cls()
		print(color.InPurple("Enter Name:"))
		fmt.Scan(&selected_item)
		for k, v := range name {
			if v == selected_item {
				found_item = k
			}
		}
		if found_item == -1 {
			println(color.InRed("User Not Found."))
		} else {
			println(color.InGreen("User Found:"))

			print(color.InCyan("Name:"))
			fmt.Println(name[found_item])

			print(color.InCyan("Birth Date:"))
			fmt.Println(bdate[found_item])

			print(color.InCyan("ID Code:"))
			fmt.Println(id_code[found_item])
			println(color.InBlue("----------------------------------"))
		}

	}

}
