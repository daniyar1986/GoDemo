package main

import "fmt"

func main() {

	strArray := []string{"Hello", "World"}
	intArray := []uint{1, 2, 3, 4, 5}
	intMake := make([]uint, 5)

	for range intMake {
		fmt.Println("Hello")
	}

	for key, value := range strArray {
		fmt.Println(key, value)
	}

	for _, value := range strArray {
		fmt.Println(value)
	}

	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	fmt.Println("last element intArray", intArray[len(intArray)-1])
	fmt.Println("first element intArray", intArray[0])
	fmt.Println("two element intArray", intArray[0:2])

}
