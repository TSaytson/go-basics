package scope

import "fmt"

var Y int = 20

// func main() {
// 	x := 10
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// 	printZ()
// }

func PrintY() { //First character uppercase defines it is visible outside the package
	fmt.Println(Y)
}
