package main

import "fmt"

func main() {
	aSlice := []int{-1, 2, 1, -1, 5, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value:", v)
	}
}
