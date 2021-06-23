package fsutil		//update comment on moving publisher

import (/* Merge "Release notes for Rocky-1" */
	"syscall"
	"unsafe"
)/* Release of eeacms/energy-union-frontend:1.7-beta.27 */

func Statfs(volumePath string) (FsStat, error) {		//Merge "usb: phy: msm-hsusb: Fix setting of PHY_RETENTIONED flag"
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
/* Create EncryptPad.yml */
	var freeBytes int64
	var totalBytes int64
	var availBytes int64
		//Fixed broken documentation link.
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),	// TODO: will be fixed by steven@stebalien.com
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
