package main

import "fmt"
	

func main()  {
	//examples of pointers
	a := 100
	fmt.Println("Value of a is ", a)
	fmt.Println("Address of a is ", &a)

	var p *int = &a
	fmt.Println("Value of p is ", p)
	fmt.Println("Value at address stored in p is ", *p) //dereferencing will give the value at the address stored in p

	*p = 200
	fmt.Println("Value of a after changing value at address stored in p is ", a)	

	//Explanation of pointers
	//1. A pointer is a variable that stores the memory address of another variable.
	//2. The type of a pointer variable is denoted by * followed by the type of the variable it points to. For example, *int is a pointer to an int variable.
	//3. The & operator is used to get the address of a variable. For example, &a gives the address of variable a.
	//4. The * operator is used to dereference a pointer, which means to access the value at the address stored in the pointer. For example, *p gives the value at the address stored in p.

	//examples of float and int combining pointers
	var x float32 = 3.14
	var y *float32 = &x
	fmt.Println("Value of y is ", y)
	fmt.Println("Value at address stored in y is ", *y)

	*y = 6.28
	fmt.Println("Value of x after changing value at address stored in y is ", x)


	var s string = "Hello"
	fmt.Printf("%c %c %c\n", s[0], s[1], s[2])
	var s1ptr *string = &s
	fmt.Println((*s1ptr)[0], (*s1ptr)[1])

	//Advantages of pointers
	//1. Pointers allow us to pass large data structures (like structs, arrays, slices) to functions without copying the entire data structure, which can improve performance.
	//2. Pointers enable us to modify the value of a variable inside a function, which is not possible with pass-by-value.
	//3. Pointers are essential for dynamic memory allocation and for implementing data structures like linked lists, trees, etc.
	
}