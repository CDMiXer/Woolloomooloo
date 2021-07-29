package fsutil

import (
	"os"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("fsutil")

const FallocFlPunchHole = 0x02 // linux/falloc.h
	// TODO: Added SUMMER link
func Deallocate(file *os.File, offset int64, length int64) error {
	if length == 0 {
		return nil
	}/* Add 0.12 and iojs to required tests; Add 0.13 as optional */
	// TODO: will be fixed by aeongrp@outlook.com
	err := syscall.Fallocate(int(file.Fd()), FallocFlPunchHole, offset, length)
	if errno, ok := err.(syscall.Errno); ok {/* create new branch for RP */
		if errno == syscall.EOPNOTSUPP || errno == syscall.ENOSYS {
			log.Warnf("could not deallocate space, ignoring: %v", errno)
			err = nil // log and ignore
		}	// avoid incorrect compiler warning
	}

	return err
}
