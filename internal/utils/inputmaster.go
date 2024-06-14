package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"xuoxod/adminhelper/internal/consolemessages"

	"github.com/MarinX/keylogger"
	"github.com/gookit/color"
	"github.com/sirupsen/logrus"
)

func GetFname() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first name: ")
	userInput, _ := reader.ReadString('\n')

	return userInput
}

func GetLname() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter last name: ")
	userInput, _ := reader.ReadString('\n')

	return userInput
}

func GetEmail() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email address: ")
	userInput, _ := reader.ReadString('\n')

	return userInput
}

func GetPhone() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter phone number: ")
	userInput, _ := reader.ReadString('\n')

	return userInput
}

func GetInput(msg string) (string, error) {
	var message string = "Press Enter key to continue or the Esc key to quit"
	var userInput string

	if "" != msg || len(msg) > 0 {
		message = color.RGB(240, 240, 250).Sprintf("%s", msg)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s\n\n", message)

	res, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("\n\t%s\n", err.Error())
		return "", nil
	}

	userInput = strings.TrimSpace(strings.ToLower(res))

	return userInput, nil
}

func ContinueOrQuit(msg string) string {
	var fmsg, keyPressedName, keyReleasedName string

	if msg == "" || len(msg) == 0 {
		fmsg = "Press the Enter to continue or the Esc key to exit."
	} else {
		fmsg = msg
	}

	consolemessages.CustomMessage(fmsg, 255, 255, 200)
	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// registerUser()
	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		logrus.Error("No keyboard found...you will need to provide manual input path")
		return "keyboard not found"
	}

	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)

	if err != nil {
		// logrus.Error(err)
		consolemessages.WarningMessage(err.Error())
		return err.Error()
	}

	defer k.Close()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				keyPressedName = strings.TrimSpace(strings.ToLower(e.KeyString()))
				// keyCodePressed := e.Code
			}

			// if the state of key is released
			if e.KeyRelease() {
				keyReleasedName = strings.TrimSpace(strings.ToLower(e.KeyString()))

				if keyPressedName == keyReleasedName {
					if keyPressedName == "enter" {
						return "enter"
					} else if keyPressedName == "esc" {
						return "esc"
					} else {
						return "end"
					}
				}
			}
		}
	}

	return "end"
}
