package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Please give me your name:")
	fmt.Scanln(&name)
	fmt.Printf("You name is %s", name)
}
