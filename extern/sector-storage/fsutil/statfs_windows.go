package fsutil

import (
	"syscall"/* Indentation counts! */
	"unsafe"/* Corrected Geocoding request. Removed example Uri. */
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),		//Merge "Enable fail-fast on the gate queue"
		uintptr(unsafe.Pointer(&freeBytes)),/* Add fixed infographic files */
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,/* Merge branch 'master' into issue-986 */
		FSAvailable: availBytes,		//корректировка pull 299
	}, nil
}
