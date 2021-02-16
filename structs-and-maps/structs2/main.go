package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(*namePointer)
	fmt.Println(namePointer)
	fmt.Println(namePointer)
	fmt.Println(&namePointer)
	fmt.Println(&namePointer)
	fmt.Println(&namePointer)
	fmt.Println(&*&namePointer)
	fmt.Println(&*&namePointer)
	fmt.Println(&*&namePointer)
	printPointer(namePointer)
	printPointer(namePointer)
	printPointer(namePointer)
}

func printPointer(np *string) {
	fmt.Println(&np)
}
