package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"xuoxod/adminhelper/internal/consolemessages"

	"github.com/PiScale/hwinfo-lib"
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

func SysInfo() (map[string]interface{}, error) {
	chassis, err := hwinfo.Get_chassis()

	if err != nil {
		consolemessages.CustomMessage(fmt.Sprintf("\nError Ocurred\n%s", err.Error()), 250, 70, 70)
		return nil, err
	}

	manu := chassis.Manufacturer
	serial := chassis.SerialNumber

	cpu, err := hwinfo.Get_cpu()

	if err != nil {
		consolemessages.CustomMessage(fmt.Sprintf("\nError Ocurred\n%s", err.Error()), 250, 70, 70)
		return nil, err
	}

	cpuModel := cpu.Model
	cpuCount := cpu.Quantity
	cpuCores := cpu.Totalcores

	nicStats, err := hwinfo.Get_nic()

	if err != nil {
		consolemessages.CustomMessage(fmt.Sprintf("\nError Ocurred\n%s", err.Error()), 250, 70, 70)
		return nil, err
	}

	ramStats, err := hwinfo.Get_ram()

	if err != nil {
		consolemessages.CustomMessage(fmt.Sprintf("\nError Ocurred\n%s", err.Error()), 250, 70, 70)
		return nil, err
	}

	// hdd, err := hwinfo.Get_hdd()
	// items := hdd.Items

	// if err != nil {
	// 	consolemessages.CustomMessage(fmt.Sprintf("\n%s", err.Error()), 240, 240, 50)
	// 	return nil
	// }

	// fmt.Println(items)

	// hddModel := hdd.Items[0].Model
	// hddSerial := hdd.Items[0].SerialNumber
	// hddFirm := hdd.Items[0].Firmware
	// hddSize := hdd.Items[0].Size
	// hddName := hdd.Items[0].DevName
	// hddBus := hdd.Items[0].Bus

	mb, err := hwinfo.Get_motherboard()

	if err != nil {
		consolemessages.CustomMessage(fmt.Sprintf("\n%s", err.Error()), 240, 240, 50)
		return nil, err
	}

	mbModel := mb.Model
	mbSerial := mb.SerialNumber

	systemInfo := fmt.Sprintf("Manufacturer:\t%s\n\t  Serial:\t%s\n\nCPU\n\tModel:\t%s\n\tCount: \t%v\n\tCores: \t%d\n\nMotherboard\n\t  Model:\t%s\n\t Serial:\t%s", manu, serial, cpuModel, cpuCount, cpuCores, mbModel, mbSerial)

	if len(nicStats.Items) > 0 {
		systemInfo += "\n\nNetwork\n"

		for _, item := range nicStats.Items {
			systemInfo += fmt.Sprintf("\tName:\t%s\n\t MAC:\t%v\n\n", item.IfName, item.MAC)
			// fmt.Printf("Name:\t%v\n\tMAC:\t%v\n", item.IfName, item.MAC)
		}
	}

	if len(ramStats.Items) > 0 {
		systemInfo += "\n\nRAM\n"

		for _, item := range ramStats.Items {
			systemInfo += fmt.Sprintf("\tVendor:\t%s\n\tSerial:\t%s\n\t  Size:\t%s\n\t  Bank:\t%s\n\t Speed:\t%s\n\n", item.Manufacturer, item.SerialNumber, item.Size, item.BankLocator, item.ClockSpeed)
		}
	}

	boxProps := make(map[string]interface{})
	boxProps["title"] = "System Info"
	boxProps["body"] = systemInfo

	return boxProps, nil
}

/*
PrintSysInfo: prints dmi information to console.
@param keyword string:

	bios
	system
	baseboard
	chassis
	processor
	memory
	cache
	connector
	slot
*/
func PrintSysInfo(keyword string) {
	if keyword == "" || len(keyword) == 0 {
		cmd := exec.Command("/usr/sbin/dmidecode")
		buf, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
		}

		output := strings.Split(string(buf), "\n")

		for _, item := range output {
			fmt.Println(item)
		}

		return
	}

	cmd := exec.Command("/usr/sbin/dmidecode", "-t", keyword)
	buf, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	output := strings.Split(string(buf), "\n")

	for _, item := range output {
		fmt.Println(item)
	}
}
