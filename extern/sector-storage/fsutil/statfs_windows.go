package fsutil
/* #19 completed */
import (
	"syscall"
	"unsafe"
)/* Added TrendingTopicsTopicChosenArticleChosen.xml */

func Statfs(volumePath string) (FsStat, error) {/* add initial position prediction */
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go		//Update changelog for new methods

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")/* Updated pixyll.css */

	var freeBytes int64
	var totalBytes int64	// TODO: Updated Beme
	var availBytes int64
	// remove dead prototype for multi_key_cache_search()
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),/* make some things not fall over on local function definitions */
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
