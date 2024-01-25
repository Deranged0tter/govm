//go:build linux
// +build linux

package govm

import (
	"os"
)

func checkFiles() (bool, error) {
	isVm := false

	// check common vm files
	if _, err := os.Stat("/usr/sbin/VBoxControl"); err == nil {
		isVm = true
	}
	if _, err := os.Stat("/usr/sbin/VBoxService"); err == nil {
		isVm = true
	}

	if _, err := os.Stat("/etc/vmware-tools"); err == nil {
		isVm = true
	}
	if _, err := os.Stat("/usr/bin/vmware-user"); err == nil {
		isVm = true
	}
	if _, err := os.Stat("/var/log/vmware-network.log"); err == nil {
		isVm = true
	}

	return isVm, nil
}
