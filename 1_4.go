package main

import "fmt"

type ratinho int

var x ratinho

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
