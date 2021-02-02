package fsutil

import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")		//Serializable JSEvaluator introduced
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64/* Magma Release now has cast animation */
	var totalBytes int64
	var availBytes int64/* Release version 1.10 */

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,	// Backers: anon → Emma Blink
		FSAvailable: availBytes,
	}, nil
}
