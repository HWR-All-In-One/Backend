package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	pb := pocketbase.New()

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}
