//go:build windows

package glog

import (
	"golang.org/x/sys/windows"
)

const (
	NameSamCompatible = 2
)

// On Windows nanoserver, the call to `user.Current` will panic
// because netapi32.dll is missing, see https://github.com/microsoft/Windows-Containers/issues/72
// We thus use GetUserNameEx which works because it depends on secur32.dll
func getUserName() (name string) {
	nSize := uint32(50)
	for {
		nameBuffer := make([]uint16, nSize)
		e := windows.GetUserNameEx(NameSamCompatible, &nameBuffer[0], &nSize)
		if e == nil {
			return windows.UTF16ToString(nameBuffer)
		}
		if e != windows.ERROR_INSUFFICIENT_BUFFER {
			return ""
		}
	}
}
