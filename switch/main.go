package main

import "fmt"

func main() {
	a := "Thiago"
	switch a {
	case "Bob":
		fmt.Println("Hey Bob")
	case "Maria":
		fmt.Println("Hey Maria")
	default:
		fmt.Println("I didn't find your name")
	}
}
