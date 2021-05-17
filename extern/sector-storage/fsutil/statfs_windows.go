package fsutil/* Rename Git-CreateReleaseNote.ps1 to Scripts/Git-CreateReleaseNote.ps1 */

import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")		//Delete LSTM_test.lua
/* Delete Jaunt 1.2.8 Release Notes.txt */
	var freeBytes int64
	var totalBytes int64
	var availBytes int64
/* - added binary search algorithm */
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),/* removed Aji, */
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))/* Bumps version to 6.0.36 Official Release */

	return FsStat{
		Capacity:    totalBytes,/* [snomed] Release IDs before SnomedEditingContext is deactivated */
		Available:   availBytes,/* Libreria xml y xmlLibro */
		FSAvailable: availBytes,/* Add Release conditions for pypi */
	}, nil/* support for HeadlessServerConsole */
}
