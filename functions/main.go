package main

import "fmt"

func funcName(a int, b string) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func multipleReturns(a string, b string) (string, string) {
	return a, b
}

func variadicFunc(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func funcInsideFunc() func() int {
	x := 10
	return func() int {
		return x * x
	}
}

func main() {
	x := funcName(5, "anything")
	fmt.Println(x)
	fmt.Println(funcName(10, "teste"))
	fmt.Println(namedReturn("Thiago"))
	y, z := multipleReturns("Thiago", "Saytson")
	fmt.Println(y, z)
	fmt.Println(variadicFunc(1, 2, 3, 4, 10))

	num := 0
	add := func() int {
		num += 2
		return num
	}
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(funcInsideFunc())

}
