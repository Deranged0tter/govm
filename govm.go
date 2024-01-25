package govm

import (
	"encoding/json"
	"errors"
	"net"
	"runtime"
	"strings"

	"github.com/pbnjay/memory"
)

// run all checks and return the results as a json
func CheckAll() string {
	isCheckMac, _ := CheckMacAdr()
	isCheckProcesses, _ := CheckProcesses()
	isCheckFiles, _ := CheckFiles()
	isCheckCores, _ := CheckCores(0)
	isCheckRam, _ := CheckRam(0)
	isCheckOnline := CheckOnline()

	results := checkAllReturn{
		CheckMac:       isCheckMac,
		CheckProcesses: isCheckProcesses,
		CheckFiles:     isCheckFiles,
		CheckCores:     isCheckCores,
		CheckRam:       isCheckRam,
		CheckOnline:    isCheckOnline,
	}

	bResults, _ := json.Marshal(results)
	var sResults string = string(bResults)

	return sResults
}

// check the machine's mac address for known hypervisor addresses
func CheckMacAdr() (bool, error) {
	// get all mac addresses
	macAddresses, err := getMacAddresses()
	if err != nil {
		return false, err
	}

	// check for known vm mac address patterns
	for _, mAdr := range macAddresses {
		if strings.Contains(mAdr, "00:50:56") || strings.Contains(mAdr, "00:0C:29") || strings.Contains(mAdr, "00:05:69") || strings.Contains(mAdr, "00:1C:14") || strings.Contains(mAdr, "08:00:27") {
			return true, nil
		}
	}

	return false, nil
}

// check running processes for processes containing "vm" or "vbox"
func CheckProcesses() (bool, error) {
	// get all processes
	processes, err := getAllProcesses()
	if err != nil {
		return false, err
	}

	// check if any processes contain "vm" or "vbox"
	for _, proc := range processes {
		if strings.Contains(strings.ToLower(proc.Executable()), "vm") || strings.Contains(strings.ToLower(proc.Executable()), "vbox") {
			return true, nil
		}
	}

	return false, nil
}

// check if files related to running a vm exist on system
func CheckFiles() (bool, error) {
	isVm, err := checkFiles()
	if err != nil {
		return false, err
	}

	return isVm, nil
}

// check whether the machine has less than or equal to x number of cores (default is 2, leave as 0 for default)
func CheckCores(count int) (bool, error) {
	if count < 0 {
		return false, errors.New("count must be positive")
	}

	cores := 2
	if count != 0 {
		cores = count
	}

	numOfProcessors := runtime.NumCPU()

	if cores <= numOfProcessors {
		return false, nil
	} else {
		return true, nil
	}
}

// check whether the machine has less than or equal to x mb of ram (default is 4196, leave as 0 for default)
func CheckRam(mb uint64) (bool, error) {
	var ram uint64 = 4196 * 1048576
	if mb != 0 {
		ram = mb * 1048576
	}

	amountOfRam := memory.TotalMemory()

	// check if ram is less than or equal to machine's ram
	if ram <= amountOfRam {
		return false, nil
	} else {
		return true, nil
	}
}

// check if machine can access 8.8.8.8
func CheckOnline() bool {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}
