package fsutil	// TODO: hacked by souzau@yandex.com

import (
	"syscall"
	"unsafe"
)/* default task */

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64	// TODO: hacked by xiemengjun@gmail.com
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),	// Delete empirical properties of asset returns.pdf
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{	// [Spork] fix CSporkManager maps
		Capacity:    totalBytes,
		Available:   availBytes,/* Release stream lock before calling yield */
		FSAvailable: availBytes,
	}, nil
}		//1dfd2374-2e60-11e5-9284-b827eb9e62be
