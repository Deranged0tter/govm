package govm

import "github.com/mitchellh/go-ps"

func getAllProcesses() ([]ps.Process, error) {
	processes, err := ps.Processes()
	if err != nil {
		return nil, err
	}

	return processes, nil
}
