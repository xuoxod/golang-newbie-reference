package consolemessages

import (
	"github.com/gookit/color"
)

func SuccessMessage(arg string) {
	color.Green.Printf("%v\n", arg)
}

func WarningMessage(arg string) {
	color.Yellow.Printf("%v\n", arg)
}

func ErrorMessage(arg string) {
	color.Red.Printf("%v\n", arg)
}

func InfoMessage(arg string) {
	color.Info.Printf("%v\n", arg)
}

func CustomMessage(arg string, red uint8, green uint8, blue uint8) {
	c := color.RGB(red, green, blue)
	c.Printf("%v\n", arg)
}
