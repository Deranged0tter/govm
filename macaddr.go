package govm

import "net"

func getMacAddresses() ([]string, error) {
	// get interfaces
	interfaces, err := getAllInterfaces()
	if err != nil {
		return nil, err
	}

	var macAdrs []string
	// loop through all interfaces
	for _, interf := range interfaces {
		macAdr := interf.HardwareAddr.String()

		if macAdr != "" {
			macAdrs = append(macAdrs, macAdr)
		}
	}

	return macAdrs, nil
}

func getAllInterfaces() ([]net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	return interfaces, nil
}
