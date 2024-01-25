package govm

type checkAllReturn struct {
	CheckMac       bool
	CheckProcesses bool
	CheckFiles     bool
	CheckCores     bool
	CheckRam       bool
	CheckOnline    bool
}
