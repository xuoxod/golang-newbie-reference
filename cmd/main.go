package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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
		Name:    "hardware",
		Aliases: []string{"hw"},
		Usage:   "display system's hardware",
		BashComplete: func(*cli.Context) {
			consolemessages.CustomMessage("\nBash complete", 240, 240, 50)
		},
		Before: func(*cli.Context) error {
			uid := os.Getuid()

			if uid != 0 {
				return errors.New("Must run this program as Admin")
			}
			return nil
		},
		After: func(*cli.Context) error {
			return nil
		},
		Action: func(cCtx *cli.Context) error {
			// fmt.Println("added task: ", cCtx.Args().First())
			var numArgs int = cCtx.NArg()

			switch numArgs {
			case 0:
				sysInfo, err := utils.SysInfo()

				if err != nil {
					fmt.Println(err.Error())
					return nil
				}

				utils.Splash(sysInfo)

			case 2:
				// arg1 := cCtx.Args().Get(0)
				// arg2 := cCtx.Args().Get(1)
				// consolemessages.CustomMessage(fmt.Sprintf("Arg 1:\t%s\nArg 2:\t%s\n", arg1, arg2), 222, 222, 222)
				// consolemessages.CustomMessage("Running Program ...", 110, 255, 110)

			default:
				// consolemessages.CustomMessage("\nToo many arguments", 240, 240, 50)
			}

			return nil
		},
		Category: "System Information",
	},
	{
		Name:    "dmi",
		Aliases: []string{"d"},
		Usage:   "print dmi info",
		Action: func(cCtx *cli.Context) error {
			// fmt.Println("completed task: ", cCtx.Args().First())
			var numArgs int = cCtx.NArg()

			switch numArgs {
			case 0:
				consolemessages.CustomMessage("\nThis command expects an argument", 255, 110, 110)

			case 1:
				arg := strings.TrimSpace(strings.ToLower(cCtx.Args().Get(0)))
				utils.PrintSysInfo(arg)

			default:
				consolemessages.CustomMessage("\nUnexpected arguments", 220, 220, 140)
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
		consolemessages.CustomMessage(fmt.Sprintf("\t%v\n", err.Error()), 250, 88, 88)
	}
}

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
