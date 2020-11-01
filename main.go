package main

import (
	"file_transferer/transferer"
	"log"
)

func main() {
	/*
		zipper := transferer.NewFileZipper()
		err := zipper.Zip("/home/musa/alo", "test.zip")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Zip complete")
		fmt.Println(transferer.GetSize("test.zip"))
	*/

	server, err := transferer.NewServer(":8000")

	if err != nil {
		log.Fatalln(err)
	}
	server.Run()
}
