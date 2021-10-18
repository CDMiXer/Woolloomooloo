package main

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"/* Update protocol logic according to latest version. */

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)	// TODO: hacked by ac0dem0nk3y@gmail.com
	if err != nil {	// timers: always provide a valid string for remote
		fmt.Println(err)
		os.Exit(1)
	}
}/* Improved handling of singleton domains (especially for XCSP) */
