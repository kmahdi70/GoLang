package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// arr := [5] string {1:"aaaa",2:"fff",3:"fgg"} 
	start := time.Now()
	for k,v:=range os.Args{
		// fmt.Println("k=",k,",v=",v)
		fmt.Printf("k=%d,v=%s(%T,%T)\n",k,v,k,v)
	}
	// fmt.Println(arr)
	duration := time.Since(start)
	fmt.Println("Time1 =",duration)
	start = time.Now()
	fmt.Println(strings.Join(os.Args,", "))
	duration = time.Since(start)
	fmt.Println("Time2 =",duration)
}
