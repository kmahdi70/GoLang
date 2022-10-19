package main

import "fmt"

func main() {

	var students [10]string

	for i := 0; i < 10; i++ {
		fmt.Println("enter student's name: ", i+1)
		fmt.Scan(&students[i])
	}

	var n int
	fmt.Println("Enter nth student:")
	fmt.Scan(&n)
	fmt.Println(n,"th student is:", students[n-1])

}
