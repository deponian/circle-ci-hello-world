package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments")
	} else {
		fmt.Printf("There is %d arguments\n", len(os.Args)-1)
	}
}
