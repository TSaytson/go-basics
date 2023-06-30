package main

import "fmt"

const (
	Aa string = "10" //visible oustide the package
	bb        = 11
)

var c, d string = "Hello", "World"
var b int = 22

func main() {
	a := 10
	b := "hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Uouuu`
	fmt.Printf("%s \n", a) //displays a warning in the terminal
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

}
