package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for k,v:=range os.Args{
		fmt.Printf("k=%d,v=%s(%T,%T)\n",k,v,k,v)
	}
	duration := time.Since(start)
	fmt.Println("Time1 =",duration)
	start = time.Now()
	fmt.Println(strings.Join(os.Args,", "))
	duration = time.Since(start)
	fmt.Println("Time2 =",duration)
}
