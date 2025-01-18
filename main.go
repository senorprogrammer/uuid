package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func main() {
	uuidCount := flag.Int("count", 1, "number of uuids to generate")

	flag.Parse()

	for i := 0; i < *uuidCount; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("Failed to generate UUID: %v", err)
		}
		fmt.Println(uuid)
	}
}
