package main

import (
	"fmt"
	"os"
	// "os"
)

func main() {
	// arr := [5] string {1:"aaaa",2:"fff",3:"fgg"} 
	for k,v:=range os.Args{
		// fmt.Println("k=",k,",v=",v)
		fmt.Printf("k=%d,v=%s(%T,%T)\n",k,v,k,v)
	}
	// fmt.Println(arr)
}
