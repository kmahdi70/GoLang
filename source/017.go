package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome")
	var flag bool = true
	var choice int
	for flag {

		fmt.Println("1)Enter Data      2)Search     3)Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			EnterData()
		case 2:
			fmt.Println("search")
		case 3:
			flag = false
		default:
			fmt.Println("Error")
		}
	}

	fmt.Println("End")

}

func EnterData() {
	var s string
	fmt.Println("enter your name:")
	fmt.Scan(&s)

	fmt.Println("enter your family:")
	fmt.Scan(&s)

}
