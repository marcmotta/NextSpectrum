// cmd/nextspectrum/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nextspectrum/internal/nextspectrum"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nextspectrum.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
