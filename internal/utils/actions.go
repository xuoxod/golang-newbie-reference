package utils

import (
	"encoding/json"
	"fmt"
	"go/types"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"xuoxod/adminhelper/internal/models"

	"github.com/Delta456/box-cli-maker/v2"
)

type Argument interface {
	int | int32 | int64 | float32 | float64 | string | types.Map | types.Array | types.Interface | types.Struct | types.TypeList | types.Named | types.Tuple | models.User
}

type Function func()

func StringNoSpaces(arg string) bool {
	return !strings.Contains(arg, " ")
}

func DateTimeStamp() string {
	// dts := fmt.Sprint("Date: ", time.Now())

	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	return fmt.Sprintf("%v %v, %v", month, day, year)
}

func DateStamp() string {
	// dts := fmt.Sprint("Date: ", time.Now())

	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	return fmt.Sprintf("%v %v %v", month, day, year)
}

func DTS() string {
	// dts := fmt.Sprint("Date: ", time.Now())
	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	hour, minute, second := time.Now().Local().Clock()

	var suffix string
	var strDay string = fmt.Sprintf("%d", day)

	if strings.HasSuffix(strDay, "1") {
		suffix = "st"
	} else if strings.HasSuffix(strDay, "2") {
		suffix = "nd"
	} else if strings.HasSuffix(strDay, "3") {
		suffix = "rd"
	} else if strings.HasSuffix(strDay, "11") || strings.HasSuffix(strDay, "12") || strings.HasSuffix(strDay, "13") {
		suffix = "th"
	} else {
		suffix = "th"
	}

	return fmt.Sprintf("%v %v%s %v %v:%v:%v", month, day, suffix, year, hour, minute, second)
}

func DS() string {
	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	var suffix string
	var strDay string = fmt.Sprintf("%d", day)

	if strings.HasSuffix(strDay, "1") {
		suffix = "st"
	} else if strings.HasSuffix(strDay, "2") {
		suffix = "nd"
	} else if strings.HasSuffix(strDay, "3") {
		suffix = "rd"
	} else {
		suffix = "th"
	}

	return fmt.Sprintf("%v %v%s %v", month, day, suffix, year)
}

func TS() string {
	hour, minute, second := time.Now().Local().Clock()
	return fmt.Sprintf("%v:%v:%v", hour, minute, second)
}

func ExitProg(exitCode int) {
	os.Exit(exitCode)
}

func Cap(arg string) string {
	var capped string

	for i, c := range strings.Split(arg, "") {
		if i == 0 {
			capped += strings.ToUpper(c)
		} else {
			capped += c
		}
	}

	return capped
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Println("Error:\t", err.Error())
	}
}

func ComAllArgs() []string {
	return os.Args
}

func CountAllArgs() int {
	return len(ComAllArgs())
}

func ComArgs() []string {
	args := []string{}

	for i, a := range os.Args {
		if i != 0 {
			args = append(args, a)
		}
	}
	return args
}

func CountArgs() int {
	return len(ComArgs())
}

func ExecuteAfterTime(seconds int, f Function) {
	duration := time.Duration(seconds) * time.Second
	timer := time.NewTimer(duration)
	<-timer.C
	f()
}

func ReadStringListToConsole(list []string) {
	for _, l := range list {
		fmt.Println(l)
	}
}

func ToString[K comparable](argument K) string {
	return strings.TrimSpace(fmt.Sprintf("%v", argument))
}

func PrettyPrint[K comparable](arg K) {
	p, x := json.MarshalIndent(arg, "", " ")

	if x != nil {
		fmt.Println("Error ", x.Error())
	}

	fmt.Printf("\n\t%v\n\n\n", string(p))
}

func Sleep(seconds int) {
	duration := time.Duration(seconds) * time.Second
	if seconds > 0 {
		time.Sleep(duration * time.Second)
	}
}

/*
	 Splash: Create a box with text and title in the console
	 @Param: Map[string]interface
	 	Map Properties:
			title: string
			titlecolor: string
			position: string
			message: string
			type: string
			boxcolor: string
			xcoord: int
			ycoord: int
			text wrap: bool
	 @return: None
*/
func Splash(boxProperties map[string]interface{}) {
	// Box properties
	var boxTitle string
	var boxTitleColor string
	var boxPosition string
	var boxBody string
	var xCoord int
	var yCoord int
	var boxType string
	var boxColor string
	var allowWrapping bool

	boxTitle = "Program"
	boxTitleColor = "White"
	boxPosition = "Top"
	boxBody = "Welcome to the program!"
	xCoord = 5
	yCoord = 5
	boxType = "Round"
	boxColor = "Cyan"
	allowWrapping = true

	for k := range boxProperties {
		key := strings.ToLower(strings.TrimSpace(k))

		if key == "title" {
			_, ok := boxProperties[key]
			if ok {
				boxTitle = fmt.Sprintf("%v", boxProperties[k])
			}
		}

		if key == "position" {
			_, ok := boxProperties[key]
			if ok {
				boxPosition = fmt.Sprintf("%v", boxProperties[k])
			}
		}

		if key == "body" {
			_, ok := boxProperties[key]
			if ok {
				boxBody = fmt.Sprintf("%v", boxProperties[k])
			}
		}

		if key == "x" || key == "xcoord" {
			_, ok := boxProperties[key]
			if ok {
				if i, err := strconv.Atoi(fmt.Sprintf("%v", boxProperties[k])); err == nil {
					xCoord = i
				}
			}
		}

		if key == "y" || key == "ycoord" {
			_, ok := boxProperties[key]
			if ok {
				if i, err := strconv.Atoi(fmt.Sprintf("%v", boxProperties[k])); err == nil {
					yCoord = i
				}
			}
		}

		if key == "type" {
			_, ok := boxProperties[key]
			if ok {
				boxType = fmt.Sprintf("%v", boxProperties[k])
			}
		}

		if key == "boxcolor" {
			_, ok := boxProperties[key]
			if ok {
				boxColor = fmt.Sprintf("%v", boxProperties[k])
			}
		}

		if key == "titlecolor" {
			_, ok := boxProperties[key]
			if ok {
				boxTitleColor = fmt.Sprintf("%v", boxProperties[k])
			}
		}

		// if key == "wrap" {
		// 	_, ok := boxProperties[key]
		// 	if ok {
		// 		allowWrapping = fmt.Sprintf("%t", boxProperties[k])
		// 	}
		// }
	}

	boxConfig := box.Config{Px: xCoord, Py: yCoord, Type: boxType, Color: boxColor, AllowWrapping: allowWrapping, TitlePos: boxPosition, TitleColor: boxTitleColor}

	Box := box.New(boxConfig)

	Box.Print(boxTitle, boxBody)

}
