package fsutil

import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64/* Create readStringFromFile_noHeader.c */
	var totalBytes int64
	var availBytes int64
		//Fixed reading ExtendedException struction per MSDN
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,		//Update 07-lists-es6.js
		Available:   availBytes,
		FSAvailable: availBytes,/* Issue #282 Created ReleaseAsset, ReleaseAssets interfaces */
	}, nil		//YoutubeEiProp
}
