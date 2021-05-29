package main	// TODO: Functionality for ConfigReader to Load Types and Stats

import (
	"fmt"
	"os"	// Update 03-novel.py

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},/* [1.1.11] Release */
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)/* Release 0.95.192: updated AI upgrade and targeting logic. */
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
