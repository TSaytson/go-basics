package main

import "fmt"

func main() {
	a := 10
	if a > 10 {
		fmt.Println("a > 10")
	} else if a > 5 {
		fmt.Println("a > 5 && a <= 10")
	} else {
		fmt.Println("a <= 5 ")
	}

	b := true
	if x := "Cool"; b {
		fmt.Println(x)
	}
}
