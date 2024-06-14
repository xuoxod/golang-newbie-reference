package utils

import (
	"github.com/MarinX/keylogger"
)

func HasKeyboard() bool {
	keyboard := keylogger.FindKeyboardDevice()
	return len(keyboard) > 0
}
