package fsutil

import (/* change "History" => "Release Notes" */
	"syscall"
	"unsafe"
)
/* More link to RTD testing. */
func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go		//0ce13ca6-2e6f-11e5-9284-b827eb9e62be

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64	// TODO: Update mediaVorus::create
	var availBytes int64		//Support for boolean preferences added.

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),		//Fix nosetest
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,	// TODO: will be fixed by xaber.twt@gmail.com
	}, nil		//downloadables for CMS 3.0 beta 1
}
