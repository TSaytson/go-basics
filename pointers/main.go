package main

import "fmt"

func addOne(num int) int {
	num = num + 1
	return num
}

func changesValueInMemoryAddress(num *int) int {
	*num = *num + 1
	return *num
}

func main() {
	b := 10
	fmt.Println(addOne(b)) //11
	fmt.Println(b)         //10 because b was passed as value

	fmt.Println(changesValueInMemoryAddress(&b))
	fmt.Println(b) //now b is 11 because was passed as reference

	x := 10

	fmt.Println(&x)

	y := &x //x memory address reference
	fmt.Println(y)
	fmt.Println(*y) //gets the value in the memory address

	*y = 20 //changes the value in the memory address pointed by y
	fmt.Println(x)

	var integerPointer *int = &x //this memory address can only contain integers
	fmt.Println(integerPointer)  //memory address of x
	fmt.Println(*integerPointer) //value inside memory address of x
	fmt.Println(&integerPointer) //memory address of integerPointer
}
