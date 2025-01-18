package main

import (
	"flag"
	"fmt"
	"log"

	"math/rand/v2"

	"github.com/fatih/color"
	"github.com/gofrs/uuid"
)

const (
	min = 32
	max = 256
)

func main() {
	uuidCount := flag.Int("count", 1, "number of uuids to generate")
	withColor := flag.Bool("color", false, "display in random color")

	flag.Parse()

	for i := 0; i < *uuidCount; i++ {
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("Failed to generate UUID: %v", err)
		}

		if !*withColor {
			fmt.Println(uuid.String())
			continue
		}

		r := rand.IntN(max-min+1) + min
		g := rand.IntN(max-min+1) + min
		b := rand.IntN(max-min+1) + min

		color.RGB(r, g, b).Println(uuid)
	}
}
