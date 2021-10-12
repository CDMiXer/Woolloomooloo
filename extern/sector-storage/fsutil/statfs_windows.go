package fsutil
/* Release of eeacms/www-devel:20.2.20 */
import (
	"syscall"/* tfGLutAdded */
	"unsafe"
)
	// Remove Archenemy Schemes from AllCardNames.txt
func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go
/* Release 0.2.0 merge back in */
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(/* Released springjdbcdao version 1.7.1 */
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),	// Adding aspectj and slf4j to archetype generated project pom
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))/* SEMPERA-2846 Release PPWCode.Vernacular.Exceptions 2.1.0. */

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
