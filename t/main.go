package main

import "fmt"

func main() {
	path := []int{1, 2}
	fmt.Println(path)
	P := path
	P[0] = 7
	fmt.Println(path)
}