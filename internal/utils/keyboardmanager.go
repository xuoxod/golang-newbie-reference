package utils

import (
	"fmt"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/sirupsen/logrus"
)

func DetectedKeyEvent() { // find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// registerUser()
	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return
	}

	logrus.Println("Found a keyboard at", keyboard)

	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)

	if err != nil {
		logrus.Error(err)
		return
	}

	defer k.Close()

	events := k.Read()
	isRoot := k.IsRoot()

	fmt.Printf("Is Root? [%v]\n\n", isRoot)

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				keyNamePressed := strings.TrimSpace(strings.ToLower(e.KeyString()))
				keyCodePressed := e.Code
				keyPressedTime := e.Time.Sec
				logrus.Println("[event] press key ", keyNamePressed)
				logrus.Println("[event] press key code ", keyCodePressed)
				logrus.Println("[event] press key time ", keyPressedTime)
			}

			// if the state of key is released
			if e.KeyRelease() {
				keyNameReleased := e.KeyString()
				keyCodeReleased := e.Code
				keyReleasedTime := e.Time.Sec
				logrus.Println("[event] release key ", keyNameReleased)
				logrus.Println("[event] release key code ", keyCodeReleased)
				logrus.Println("[event] release key time ", keyReleasedTime)
			}
		}
	}
}
