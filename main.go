package main

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func main() {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}
	fmt.Println(uuid)
}
