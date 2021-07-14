package main

import (
	"fmt"
	"os"
/* Update activity_activity_reporte.xml */
	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)		//Create singles.py
	if err != nil {/* Correct service */
		fmt.Println(err)
		os.Exit(1)
	}
}
