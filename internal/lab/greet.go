package lab

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Greet() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(cCtx *cli.Context) error {
			fmt.Printf("Hello %q!\n\n", cCtx.Args().Get(0))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
