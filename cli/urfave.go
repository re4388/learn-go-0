package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Run_cli0() {
	app := &cli.App{
		Name:  "GCP Account Switcher",
		Usage: "Choose GCP Account",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
