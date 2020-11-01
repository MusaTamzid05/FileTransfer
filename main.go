package main

import (
	"file_transferer/transferer"
	"flag"
	"fmt"
	"log"
)

func RunServer(address string) {

	server, err := transferer.NewServer(":8000")

	if err != nil {
		log.Fatalln(err)
	}
	server.Run()
}

func RunClient(address, path string) {

	client, err := transferer.NewClient(":8000")

	if err != nil {
		log.Fatalln(err)
	}
	client.Run(path)
}

func ZipperExample() {

	zipper := transferer.NewFileZipper()
	err := zipper.Zip("/home/musa/alo", "test.zip")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Zip complete")
	fmt.Println(transferer.GetSize("test.zip"))
}

func Usage() string {
	return "serverAddr server_address (serverPtr) path  path_to_dst"
}

func main() {
	serverAddressPtr := flag.String("serverAddr", "", "The main server")
	serverPtr := flag.Bool("server", false, "True to run as server")
	pathPtr := flag.String("path", "", "The path to the file to transfer")

	flag.Parse()

	if *serverAddressPtr == "" {
		log.Fatalln(Usage())
	}

	if *serverPtr {
		log.Println("Running as server")
		RunServer(*serverAddressPtr)
		return
	}

	log.Println("Running as client")
	RunClient(*serverAddressPtr, *pathPtr)

}
