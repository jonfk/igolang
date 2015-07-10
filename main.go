package main

import (
	"github.com/jonfk/igolang/igolang"
	"log"
	"os"
)

func main() {
	connFilePath := os.Args[1]
	log.Printf("connection_file path: %s", connFilePath)

	kernel := igolang.NewKernel(connFilePath)
	log.Println("new Kernel created")
	kernel.Run()
}
