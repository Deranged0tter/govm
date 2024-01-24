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

	return isVm, nil
}
