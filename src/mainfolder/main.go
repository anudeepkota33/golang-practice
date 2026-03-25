package main

import "fmt"
	

func main()  {
	fmt.Println("Hello World")
	
	//types of declarations
	//1. var
	//2. const
	//3. type
	//4. func
	
	// var a int = 23
	// var b int
	// var c = 100
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// var s rune = 'A'
	// fmt.Println(s)

	// fmt.Printf("%d, %[1]c\n", s)

	// var e, f, g int = 1, 2, 3
	// fmt.Println(e, f, g)

	// //shortcut way of declaring and initializing a variable only inside a function
	// h := 1000
	// fmt.Println(h)

	// //lower case functions and variables are not exported outside the package
	// //upper case functions and variables are exported outside the package when the package is imported

	// j, k, l := 3, 4, 5
	// fmt.Println(j, k, l)

	// //atleast 1 new variable should be there in the short declaration
	// //j, k, l := 1, 2, 3 //error: no new variables on left side of :=
	// n, k, l:= 1, 2, 3
	// fmt.Println(n, k, l)

	//untyped
	// const pi = 3.14
	
	// var x int = 100
	// var y float32 = float32(x)
	// fmt.Println(y)

	// //typed
	// const pi1 float32 = 3.14

	// var x1 float32 = pi1
	// var y1 float32 = pi1
	// println(x1, y1)


	// type apples int
	// var kashmirrappes apples
	// var washingtonapples apples
	
	// a, b := Helloworld(10, 5)
	// fmt.Println(a, b)

	// var x1 int16 = 100
	// var x2 int32 = 200

	// var x3 int32 = int32(x1) + x2
	// fmt.Println(x3)
	// fmt.Printf("%d %[1]x\n", x3)

	// var x4 float32 = 20.34
	// fmt.Printf("%f %[1]e %[1]g\n", x4)

	// var b bool
	// b = true
	// fmt.Printf("%t\n", b)
	
	var s string = "Golang \n Programming"
	fmt.Printf("%s\n", s)
	fmt.Printf("%c\n", s[0])
	fmt.Printf("%s\n", s[0:6])

	var s1 string = `Golang \n Programming`
	fmt.Printf("%s\n", s1)

	var a1 rune = 'C'
	fmt.Printf("%c %[1]q\n", a1)

	s2 := []rune(s)
	s2[0] = 'P'
	fmt.Printf("%s\n", string(s2))

}

func Helloworld(x int, y int)(z int, s int)  {
	z = x + y
	s = x - y
	return 
}