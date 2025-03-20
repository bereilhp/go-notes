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
	var test string
	test = "John"
	var y, z int = 1, 3
	hello, world := 2, "World!"

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
	fmt.Printf("variable test=%v is of type %T \n", test, test)
	fmt.Printf("variable y=%v is of type %T \n", y, y)
	fmt.Printf("variable z=%v is of type %T \n", z, z)
	fmt.Printf("variable hello=%v is of type %T \n", hello, hello)
	fmt.Printf("variable world=%v is of type %T \n", world, world)

}
