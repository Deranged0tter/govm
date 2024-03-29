//go:build windows
// +build windows

package govm

import "os"

var files []string = []string{"C:\\windows\\System32\\Drivers\\Vmmouse.sys", "C:\\windows\\System32\\Drivers\\vm3dgl.dll", "C:\\windows\\System32\\Drivers\\vmdum.dll", "C:\\windows\\System32\\Drivers\\vm3dver.dll", "C:\\windows\\System32\\Drivers\\vmtray.dll", "C:\\windows\\System32\\Drivers\\VMToolsHook.dll", "C:\\windows\\System32\\Drivers\\vmmousever.dll", "C:\\windows\\System32\\Drivers\\vmhgfs.dll", "C:\\windows\\System32\\Drivers\\vmGuestLib.dll", "C:\\windows\\System32\\Drivers\\VmGuestLibJava.dll", "C:\\windows\\System32\\Driversvmhgfs.dll", "C:\\windows\\System32\\Drivers\\VBoxMouse.sys", "C:\\windows\\System32\\Drivers\\VBoxGuest.sys", "C:\\windows\\System32\\Drivers\\VBoxSF.sys", "C:\\windows\\System32\\Drivers\\VBoxVideo.sys", "C:\\windows\\System32\\vboxdisp.dll", "C:\\windows\\System32\\vboxhook.dll", "C:\\windows\\System32\\vboxmrxnp.dll", "C:\\windows\\System32\\vboxogl.dll", "C:\\windows\\System32\\vboxoglarrayspu.dll", "C:\\windows\\System32\\vboxoglcrutil.dll", "C:\\windows\\System32\\vboxoglerrorspu.dll", "C:\\windows\\System32\\vboxoglfeedbackspu.dll", "C:\\windows\\System32\\vboxoglpackspu.dll", "C:\\windows\\System32\\vboxoglpassthroughspu.dll", "C:\\windows\\System32\\vboxservice.exe", "C:\\windows\\System32\\vboxtray.exe", "C:\\windows\\System32\\VBoxControl.exe"}

func checkFiles() (bool, error) {
	isVm := false

	for _, path := range files {
		if _, err := os.Stat(path); err == nil {
			isVm = true
			return isVm, nil
		}
	}

	return false, nil
}
