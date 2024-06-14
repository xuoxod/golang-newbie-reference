package main

import (
	"fmt"
	"os"
	"xuoxod/adminhelper/internal/consolemessages"
	"xuoxod/adminhelper/internal/utils"

	"github.com/urfave/cli/v2"
)

var InputErrors map[string]bool = make(map[string]bool)
var InternalErrorCount int = 0
var flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:     "english",
		Value:    "english",
		Usage:    "language for the app",
		Aliases:  []string{"en", "enu"},
		Category: "Language to use",
	},
	&cli.StringFlag{
		Name:     "chinese",
		Value:    "chinese",
		Usage:    "language for the app",
		Aliases:  []string{"zh", "ch"},
		Category: "Language to use:",
	},
	&cli.StringFlag{
		Name:     "account",
		Value:    "user account",
		Usage:    "do something to the user's account",
		Aliases:  []string{"usr", "u"},
		Category: "Account activity:",
	},
	&cli.StringFlag{
		Name:     "quiz",
		Value:    "user account",
		Usage:    "test the app's functionality",
		Aliases:  []string{"test", "t"},
		Category: "Program Tests:",
	},
}
var commands []*cli.Command = []*cli.Command{
	{
		Name:    "groups",
		Aliases: []string{"g"},
		Usage:   "list a user's groups",
		BashComplete: func(*cli.Context) {
			consolemessages.CustomMessage("\nBash complete", 240, 240, 50)
		},
		Before: func(*cli.Context) error {
			boxProps := make(map[string]interface{})
			boxProps["title"] = "Groups Command"
			boxProps["body"] = "This function is called before"
			utils.Splash(boxProps)
			// consolemessages.CustomMessage("\nBefore", 240, 240, 50)
			return nil
		},
		After: func(*cli.Context) error {
			consolemessages.CustomMessage("\nAfter", 240, 240, 50)
			return nil
		},
		Action: func(cCtx *cli.Context) error {
			// fmt.Println("added task: ", cCtx.Args().First())
			var numArgs int = cCtx.NArg()

			switch numArgs {
			case 0:
				consolemessages.CustomMessage("\nMissing arguments", 255, 110, 110)

			case 2:
				arg1 := cCtx.Args().Get(0)
				arg2 := cCtx.Args().Get(1)
				consolemessages.CustomMessage(fmt.Sprintf("Arg 1:\t%s\nArg 2:\t%s\n", arg1, arg2), 222, 222, 222)
				consolemessages.CustomMessage("Running Program ...", 110, 255, 110)

			default:
				consolemessages.CustomMessage("\nToo many arguments", 240, 240, 50)
			}

			return nil
		},
		Category: "User account activity:",
	},
	{
		Name:    "lock",
		Aliases: []string{"l"},
		Usage:   "lock user's account",
		Action: func(cCtx *cli.Context) error {
			// fmt.Println("completed task: ", cCtx.Args().First())
			var numArgs int = cCtx.NArg()

			switch numArgs {
			case 0:
				consolemessages.CustomMessage("\nMissing arguments", 255, 110, 110)

			case 1:
				arg := cCtx.Args().Get(0)
				consolemessages.CustomMessage(fmt.Sprintf("Arg :\t%s\n", arg), 222, 222, 222)
				consolemessages.CustomMessage("Running Program ...", 110, 255, 110)

			default:
				consolemessages.CustomMessage("\nAn error ocurred", 220, 220, 140)
			}
			return nil
		},
		Category: "User account activity:",
	},
	{
		Name:    "template",
		Aliases: []string{"t"},
		Usage:   "options for task templates",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("new user account: ", cCtx.Args().First())
					return nil
				},
				Category: "Account activity:",
			},
			{
				Name:  "remove",
				Usage: "remove an existing template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("removed user's account: ", cCtx.Args().First())
					return nil
				},
				Category: "User account activity:",
			},
		},
	},
}

const InternalErrorLimit int = 2

func main() {
	utils.ClearScreen()
	// consolemessages.CustomMessage("\t   Starting Up ...\n\n", 190, 250, 190)
	// Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	// Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

	// splash()

	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "Admin Assistant",
		Usage:                "Help the admin do ...",
		Flags:                flags,
		Commands:             commands,
		Action:               actionHandler,
		Version:              "0.0.1",
	}

	if err := app.Run(os.Args); err != nil {
		// log.Fatal(err)
		consolemessages.CustomMessage("Error happend at program start", 190, 250, 180)
		consolemessages.CustomMessage(fmt.Sprintf("\t%v\n\n", err.Error()), 240, 100, 155)
	}
}

/* func splash(title, body string) {
	boxConfig := box.Config{Px: 5, Py: 2, Type: "Round", Color: "Cyan", AllowWrapping: true, TitlePos: "Top", TitleColor: "White"}

	Box := box.New(boxConfig)

	Box.Print(title, body)
} */

func actionHandler(cCtx *cli.Context) error {
	cmdArgs := cCtx.NArg()
	cmdFlags := cCtx.NumFlags()
	command := cCtx.Command

	fmt.Printf("Command:\t%v\n", command)

	if cmdArgs > 0 {
		strArgs := fmt.Sprintf("Number of Arguments:\t%d", cmdArgs)
		fmt.Println(strArgs)
	}

	if cmdFlags > 0 {
		strFlags := fmt.Sprintf("Number of Flags:\t%d", cmdFlags)

		fmt.Println(strFlags)
	}

	return nil
}
