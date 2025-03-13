package main

import (
	"fmt"
)

func main() {
	var student1 string = "John"
	var student2 = "Jane"
	x := 2
	var a int
	var b bool
	var c string
	var count int = 42
	var message string = "go find type"
	var isCheck bool = true
	var amount float32 = 10.2

	fmt.Printf("variable count=%v is of type %T \n", count, count)
	fmt.Printf("variable message='%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck='%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount=%v is of type %T \n", amount, amount)
	fmt.Printf("variable student1=%v is of type %T \n", student1, student1)
	fmt.Printf("variable student2=%v is of type %T \n", student2, student2)
	fmt.Printf("variable x=%v is of type %T \n", x, x)
	fmt.Printf("variable a=%v is of type %T \n", a, a)
	fmt.Printf("variable b=%v is of type %T \n", b, b)
	fmt.Printf("variable c=%v is of type %T \n", c, c)
}
