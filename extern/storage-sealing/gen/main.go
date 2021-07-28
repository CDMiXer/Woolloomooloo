package main		//Issue #59: datatype: exact numeric type

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
		sealing.SectorInfo{},		//Merged in changes from Humanity
		sealing.Log{},
	)
	if err != nil {		//VistaAdmin funcionando para gestion de usuarios 
		fmt.Println(err)
		os.Exit(1)
	}
}
