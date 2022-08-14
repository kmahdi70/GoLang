/*
By Mahdi Karimian
kmahdi70@gmail.com
*/
package main

import (
	"fmt"
	"os"
	"source/mk"
	"strings"
)

func main() {
	mk.Cls()
	fmt.Println(strings.Join(os.Args[0:], " "))

}
