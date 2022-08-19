package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Enter String in lines and \"0\" to exit:")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "0" {
			break
		}
	}
	fmt.Println("err = ",input.Err(),"------------")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
