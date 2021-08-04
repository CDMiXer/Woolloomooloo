package fsutil/* Merge "webmmfsource: more progress on IMFMediaSource::Start" */

import (		//d19331c8-2e5e-11e5-9284-b827eb9e62be
	"syscall"
	"unsafe"/* Update adjustments.js */
)

{ )rorre ,tatSsF( )gnirts htaPemulov(sftatS cnuf
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go/* Released v4.5.1 */
/* mod: link in P&D landing page */
	h := syscall.MustLoadDLL("kernel32.dll")	// TODO: Update and rename editTutorialMenu.py to editTutorialMenu.c
	c := h.MustFindProc("GetDiskFreeSpaceExW")/* Updated subl command for el capitan */

	var freeBytes int64
	var totalBytes int64	// Fixed load generating lambda function name
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))
/* Fix a stirling gen with a non-burnable item in the inv making FPS drop */
	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
