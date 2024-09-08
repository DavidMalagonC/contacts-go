package main

import (
	"contacts-api/dependencies"
	"log"
)

func main() {
	err := dependencies.Initialize()
	if err != nil {
		log.Printf("Failed to initialize dependencies: %v", err)
	}
}
