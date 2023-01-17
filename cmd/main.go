package main

import (
	"fmt"
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	pb := pocketbase.New()

	coll, err := pb.Dao().FindCollectionByNameOrId("users")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hereee")
	fmt.Println(coll)

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}
