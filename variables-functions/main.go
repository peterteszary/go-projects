package main

import (
	"fmt"
	"log"
)

var s = "seven"

func main() {

	log.Println(s)

	fmt.Println("Hello World.")

	var whatToSay string 
	var i int

	whatToSay = "Goodby, cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOherThingThaWasSaid := saySometing()

	fmt.Println("The function returned", whatWasSaid, theOherThingThaWasSaid)
}

func saySometing() (string, string) {
	return "something", "else"

}