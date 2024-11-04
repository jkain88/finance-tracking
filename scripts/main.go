package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <command>", os.Args[0])
	}
	command := os.Args[1]

	switch command {
	case "populate":
		PopulateTransactions()
	case "hash":
		hash, err := HashPassword(os.Args[2])
		if err != nil {
			log.Fatal("Hash error")
			return
		}
		fmt.Println(hash)
	default:
		log.Fatal("Unkown command")
	}

}
