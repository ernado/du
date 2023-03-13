//go:build windows

package du

import (
	"errors"
	"syscall"
	"unsafe"
)

func getUsage(path string) (*Info, error) {
	i := Info{}
	h, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return nil, err
	}
	c, err := h.FindProc("GetDiskFreeSpaceExW")
	if err != nil {
		return nil, err
	}
	_, _, err = c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&i.Free)),
		uintptr(unsafe.Pointer(&i.Total)),
		uintptr(unsafe.Pointer(&i.Available)),
	)
	ok := syscall.Errno(0)
	if errors.Is(err, ok) {
		err = nil
	}
	if err != nil {
		return nil, err
	}
	return &i, nil
}
