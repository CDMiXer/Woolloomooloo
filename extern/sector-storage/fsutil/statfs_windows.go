package fsutil

import (
	"syscall"
	"unsafe"
)		//Changed particle hook to be used with IMetaIconProvider

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go	// TODO: hacked by jon@atack.com

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")/* Relaunched Travis CI notification */

	var freeBytes int64
	var totalBytes int64		//Add new interface IFileConfiguration.
	var availBytes int64
/* detail pane reworked */
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{/* Create pbsort */
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
