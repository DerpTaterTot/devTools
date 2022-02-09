package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DerpTaterTot/devTools/tools"
	"github.com/urfave/cli/v2"
)

func main() {
	var rawtext string
	var shift int

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "ceasarencrypt",
				Aliases: []string{"a"},
				Usage:   "encrypts a given string",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "rawtext",
						Aliases:     []string{"rt"},
						Usage:       "raw text to be encrypted",
						Destination: &rawtext,
					},
					&cli.IntFlag{
						Name:        "shift",
						Aliases:     []string{"s"},
						Usage:       "shift in the encyrption",
						Destination: &shift,
					},
				},
				Action: func(c *cli.Context) error {
					encryptedtext := tools.CeasarEncryption(shift, rawtext)

					fmt.Printf("Encrypted text with a Ceasar cipher: %s", encryptedtext)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
