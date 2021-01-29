package fsutil/* Release details for Launcher 0.44 */
	// TODO: DOC: Added requirements file
import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64	// TODO: hacked by alex.gaynor@gmail.com
	var availBytes int64		//Use only artifactId for unique identifier

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),/* Release 0.95.194: Crash fix */
		uintptr(unsafe.Pointer(&freeBytes)),/* Update meso.py */
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,/* Corrected spelling mistake in sbt.bat */
		FSAvailable: availBytes,
	}, nil
}
