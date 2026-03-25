package main

import "fmt"

func main() {
	//examples of functions call by value and call by reference
	a := 10
	b := 20
	fmt.Println("Before calling function, value of a is ", a, " and value of b is ", b)
	sum, diff := Helloworld(a, b)
	fmt.Println("After calling function, sum is ", sum, " and difference is ", diff)
	fmt.Println("After calling function, value of a is ", a, " and value of b is ", b)

	//Explanation of call by value and call by reference
	//1. In call by value, a copy of the actual parameter is passed to the function. This means that any changes made to the parameter inside the function do not affect the original variable outside the function. In the above example, a and b are passed by value to the Helloworld function.
	//2. In call by reference, a reference (or pointer) to the actual parameter is passed to the function. This means that any changes made to the parameter inside the function will affect the original variable outside the function. In Go, we can achieve call by reference using pointers.

	//Example of call by reference
	aPtr := &a
	bPtr := &b
	fmt.Println("Before calling function, value of a is ", a, " and value of b is ", b)
	sum, diff = HelloworldRef(aPtr, bPtr)
	fmt.Println("After calling function, sum is ", sum, " and difference is ", diff)
	fmt.Println("After calling function, value of a is ", a, " and value of b is ", b)

	//veriadic functions
	fmt.Println("Sum of 1, 2, 3 is ", sumVariadic(1, 2, 3))
	fmt.Println("Sum of 4, 5, 6, 7 is ", sumVariadic(4, 5, 6, 7))
}

func sumVariadic(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Helloworld(x int, y int) (z int, s int) {
	z = x + y
	s = x - y
	return
}

func HelloworldRef(xPtr *int, yPtr *int) (z int, s int) {
	z = *xPtr + *yPtr
	s = *xPtr - *yPtr
	return
}
