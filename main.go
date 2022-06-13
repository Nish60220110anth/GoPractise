package main

import (
	"fmt"
	"os"
)

func main() {
	// Read from external file
	fileName := os.Args[1]

	fmt.Println(fmt.Scanf("File : %v ", fileName))

}
