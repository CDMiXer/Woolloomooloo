package main

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"	// comment new changes

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",		//annotations for tucanoan|perry_priest1985
		sealing.Piece{},	// TODO: hacked by peterke@gmail.com
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
