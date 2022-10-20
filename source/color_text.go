package main

import "github.com/TwiN/go-color"

func main() {
    println(color.InRed(color.InBold("This is bold")))
    println(color.InRed("This is red"))
    println(color.InGreen("This is green"))
    println(color.InYellow("This is yellow"))
    println(color.InBlue("This is blue"))
    println(color.InPurple("This is purple"))
    println(color.InCyan("This is cyan"))
    println(color.InGray("This is gray"))
    println(color.InWhite("This is white"))
}