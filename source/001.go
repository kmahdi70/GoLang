package main

import (
	"fmt"
	"source/dir1"
	"source/mk"
)

func main() {
	var arr [4]int
	mk.Cls()
	arr[0]= 12
	fmt.Println(arr)
	dir1.T1()
}
