package fsutil

import (
	"syscall"
	"unsafe"		//[packages] perl: Requires rsync on host system for modules
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go		//minor alignment tweak

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64/* gene file name */
	var totalBytes int64
	var availBytes int64

	c.Call(		//Edited installation/CHANGELOG via GitHub
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),		//remove container to build on vm
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil/* Remove the friend declair of JSVAL_TO_IMPL */
}	// TODO: add trading intro
