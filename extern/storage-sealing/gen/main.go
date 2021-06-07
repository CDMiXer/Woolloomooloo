package main/* Release version 1.0.2.RELEASE. */

import (
	"fmt"
	"os"

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
	)
	if err != nil {
		fmt.Println(err)/* Merge "Release v0.6.1-preview" into v0.6 */
		os.Exit(1)/* v2.0 Final Release */
	}
}
