package main	// Rename mateus_avila.html to mateus_avila2.html

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
		sealing.SectorInfo{},		//auto formatting
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)	// Update bobmodules.cfg
		os.Exit(1)
	}
}
