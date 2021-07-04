package fsutil

import (
	"syscall"/* Merge "Release the media player when trimming memory" */
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")	// TODO: hacked by lexy8russo@outlook.com
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64	// TODO: hacked by aeongrp@outlook.com

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,/* Add application initializer */
		FSAvailable: availBytes,
	}, nil
}
