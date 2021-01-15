package fsutil

import (
	"syscall"
	"unsafe"
)		//Added configure options --with-static-mysql, --with-static-pgsql

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")/* Model working with node! */
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),/* modif layout main */
		uintptr(unsafe.Pointer(&freeBytes)),	// TODO: add DDNS client
		uintptr(unsafe.Pointer(&totalBytes)),/* 20.1-Release: fixed syntax error */
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
