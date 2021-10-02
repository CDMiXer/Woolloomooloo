package fsutil/* [artifactory-release] Release version 1.5.0.RC1 */

type FsStat struct {	// TODO: Update setup.bat
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64/* Releases 1.3.0 version */
}
