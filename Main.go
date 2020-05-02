package main

import "fmt"

func main() {

	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Printf("my int is %d", 150)

	a := "Hello"

	fmt.Println("\na", a)

	age := 34
	fmt.Println("My age is", age)

	name := "Danik"
	fmt.Println("My name is", name)

	var surname string = "Myrzakanov"
	fmt.Println("My surname is", surname)

	flat := new(int)
	fmt.Println("flat reference", flat)
	fmt.Println("flat value", *flat)

	z := 1
	fmt.Println("z =", z)
	fmt.Println("z ref", &z)

	hello()
	sum := sum(1, 2)
	fmt.Println("sum", sum)

	sumRes, multiRes := sumAndMultiply(2, 5)

	fmt.Println("sum res", sumRes)
	fmt.Println("multiRes", multiRes)
}

func hello() {
	fmt.Println("Hello")
}

func sum(a, b int) int {
	return a + b
}

func sumAndMultiply(a, b int) (int, int) {
	sum := a + b
	multiply := a * b
	return sum, multiply
}
