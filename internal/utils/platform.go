package utils

import (
	"fmt"
	"runtime"
)

func GetOS() string {
	return fmt.Sprintf("%v", runtime.GOOS)
}

func GetArch() string {
	return fmt.Sprintf("%v", runtime.GOARCH)
}

func GetRoot() string {
	return fmt.Sprintf("%v", runtime.GOROOT())
}

func GetCaller() (string, string, string) {
	pc, file, line, ok := runtime.Caller(0)

	if !ok {
		fmt.Println("Could not get the runtime caller")
		return "", "", ""
	}
	return fmt.Sprintf("%d", pc), fmt.Sprintf("%v", file), fmt.Sprintf("%d", line)

}

func CountCPU() string {
	return fmt.Sprintf("%v", runtime.NumCPU())
}
