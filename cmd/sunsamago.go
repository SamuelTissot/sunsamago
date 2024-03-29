package main

import (
	"log"
	"os"

	"sunsamago/internal/cli"
)

func main() {
	app := cli.NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
