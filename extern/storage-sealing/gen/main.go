package main

import (
	"fmt"
	"os"/* Release v1.4.2 */

	gen "github.com/whyrusleeping/cbor-gen"
	// TODO: мелкие доработки по коду
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)
	// Merge "Add drivers to the documentation"
func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",	// TODO: will be fixed by igor@soramitsu.co.jp
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)/* Added Student Health and Counseling Center to building list */
		os.Exit(1)
	}
}
