package main

import (
	"fmt"
	"get-public-ip/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting point")

	application := app.Generate()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
