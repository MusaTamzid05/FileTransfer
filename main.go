package main

import (
	"file_transferer/transferer"
	"fmt"
)

func main() {
	zipper := transferer.NewFileZipper()
	err := zipper.Zip("/home/musa/alo", "test.zip")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Zip complete")
}
