package main

import (
	"fmt"
	"os"
	"structs"
	"/src/"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: invalid number of arguments: please enter 'go run ./cmd/ filename.txt'")
		return
	}
	Colonie, err := src.HandleFormat(args[0])
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	
}
