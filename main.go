package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DerpTaterTot/devTools/tools"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "ceasarencrypt",
				Aliases: []string{"ce"},
				Usage:   "encrypts a given string with a ceasar cipher",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "rawtext",
						Aliases:  []string{"t"},
						Usage:    "raw text to be encrypted",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "shift",
						Aliases:  []string{"s"},
						Usage:    "shift in the encryption",
						Value:    0,
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					encryptedtext := tools.CeasarEncryption(c.Int("shift"), c.String("rawtext"))

					fmt.Printf("Encrypted text with a %v shift (Ceasar): %s", c.Int("shift"), encryptedtext)
					return nil
				},
			},
			{
				Name:    "ceasardecrypt",
				Aliases: []string{"cd"},
				Usage:   "decrypts a given string with a ceasar cipher",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "encryptedtext",
						Aliases:  []string{"t"},
						Usage:    "encrypted text to be decrypted",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "shift",
						Aliases: []string{"s"},
						Usage:   "shift in the decryption",
						Value:   0,
					},
					&cli.BoolFlag{
						Name:    "all",
						Aliases: []string{"a"},
						Usage:   "whether to display all decryptions (used when you don't know the shift)",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("all") {
						rawtextarray := tools.CeasarDecryptionWithoutShift(c.String("encryptedtext"))

						for i := 1; i < len(rawtextarray)-1; i++ { // ignore 0 case and final case because that's just the inital word
							fmt.Printf("Decrypted text with a %v shift (Ceasar): %s\n", i, rawtextarray[i])
						}
					} else {
						rawtext := tools.CeasarDecryptionWithShift(c.Int("shift"), c.String("encryptedtext"))

						fmt.Printf("Decrypted text with a %v shift (Ceasar): %s", c.Int("shift"), rawtext)
					}
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
