package fsutil
	// TODO: hacked by xiemengjun@gmail.com
type FsStat struct {
	Capacity    int64		//Added new_pod_repository for PodToBUILD
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// Add configuration for Clock. "java" cron does not work for now
	Reserved    int64
	// TODO: 0.4 is out.
	// non-zero when storage has configured MaxStorage	// shm/shm: add method CalcMemoryMapSize()
	Max  int64
	Used int64
}
