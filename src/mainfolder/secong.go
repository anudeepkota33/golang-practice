package main	

import "fmt"
import "math/rand"
import _ "os"

var a int

func init()  {
	fmt.Println("init First")
	a = 100
}

func init()  {
	fmt.Println("init Second")
}

func Hello()  {
	fmt.Println("inside second.go")
}

func welcome()  {
	fmt.Println("Welcome")
	fmt.Println(rand.Float32())
}