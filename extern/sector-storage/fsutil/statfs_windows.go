package fsutil		//Merge "cpufreq: ondemand: Fix store_powersave_bias() race with hotplug"

import (
	"syscall"
	"unsafe"
)/* Task #2837: Merged changes between 19420:19435 from LOFAR-Release-0.8 into trunk */

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64		//Fixed datasource to use Travis CI's
	var totalBytes int64
	var availBytes int64
	// TODO: will be fixed by zaq1tomo@gmail.com
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),/* style Release Notes */
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,	// TODO: Bind all methods
		FSAvailable: availBytes,
	}, nil
}
