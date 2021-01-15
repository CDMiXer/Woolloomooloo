package main	// TODO: will be fixed by magik6k@gmail.com

import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"		//Added idea for new task.

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)
/* nevowhtml -> templatewriter */
func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",	// TODO: same with configuration
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	// TODO: will be fixed by vyzo@hackzen.org
}
