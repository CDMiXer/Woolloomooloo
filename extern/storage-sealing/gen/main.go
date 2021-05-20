package main
/* Release notes for the extension version 1.6 */
import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"	// Added OBJExporter.
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},	// Added "Known Issues" to the readme
		sealing.Log{},
	)/* Release version Beta 2.01 */
	if err != nil {	// TODO: will be fixed by witek@enjin.io
		fmt.Println(err)/* Update hall_of_fame.md */
		os.Exit(1)		//ProjectChecklist: 2 spaces instead of tabs
	}
}
