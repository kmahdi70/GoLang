package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)
	var employee1 = map[string]int{}
	employee1 = map[string]int{"aa": 1, "bb": 3}
	fmt.Println("emp1=", employee1)
	var employee = make(map[string]string)
	employee["family"] = "rezaei"
	employee["name"] = "ali"
	fmt.Println(employee)
	fmt.Println("Len=", len(employee))
	delete(employee, "name")
	fmt.Println(employee)
	fmt.Println("Len=", len(employee))
	var i2s = map[int]string{100: "ali", 20: "ahmadi"}
	fmt.Println(i2s)

	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
}
