package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0
	for x < 10 { //while doesnt exists in go
		fmt.Println(x)
		x++
	}

	for { // infinite loop in go
		println(x)
		x++
		if x == 50 {
			continue
		}
		if x == 100 {
			break
		}
	}
}
