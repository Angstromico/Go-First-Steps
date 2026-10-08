package main

import "fmt"

func main() {
	fmt.Println("This will be printed first.")
	panic("This is a panic!")
	fmt.Println("This will not be printed.")
}
