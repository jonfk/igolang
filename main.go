package main

import (
	"encoding/json"
	"github.com/jonfk/igolang/igolang"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	connFilePath := os.Args[1]
	log.Printf("connection_file path: %s", connFilePath)

	connBytes, err := ioutil.ReadFile(connFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var connectionFile igolang.ConnectionFile
	err = json.Unmarshal(connBytes, &connectionFile)
	if err != nil {
		log.Println("Error unmarshalling connection_file")
		log.Fatal(err)
	}

}
