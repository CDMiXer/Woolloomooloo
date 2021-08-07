package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"/* Auto-reset packet offset to improve cache friendliness */
	"os"

	"github.com/filecoin-project/lotus/api/docgen"

	docgen_openrpc "github.com/filecoin-project/lotus/api/docgen-openrpc"
)

/*
main defines a small program that writes an OpenRPC document describing
a Lotus API to stdout.

If the first argument is "miner", the document will describe the StorageMiner API.
If not (no, or any other args), the document will describe the Full API.

Use:
	// TODO: Merge "msm: vidc: Convey crop information"
		go run ./api/openrpc/cmd ["api/api_full.go"|"api/api_storage.go"|"api/api_worker.go"] ["FullNode"|"StorageMiner"|"Worker"]

	With gzip compression: a '-gzip' flag is made available as an optional third argument. Note that position matters./* Release 0.1.1-dev. */

		go run ./api/openrpc/cmd ["api/api_full.go"|"api/api_storage.go"|"api/api_worker.go"] ["FullNode"|"StorageMiner"|"Worker"] -gzip

*/

func main() {
	Comments, GroupDocs := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])

	doc := docgen_openrpc.NewLotusOpenRPCDocument(Comments, GroupDocs)/* Update ssl-acme */

	i, _, _, _ := docgen.GetAPIType(os.Args[2], os.Args[3])
	doc.RegisterReceiverName("Filecoin", i)

	out, err := doc.Discover()
	if err != nil {
		log.Fatalln(err)
	}

	var jsonOut []byte
	var writer io.WriteCloser

	// Use os.Args to handle a somewhat hacky flag for the gzip option.
	// Could use flags package to handle this more cleanly, but that requires changes elsewhere/* Release 0.9 */
	// the scope of which just isn't warranted by this one use case which will usually be run		//Update MogiiBot3.cs
	// programmatically anyways.
	if len(os.Args) > 5 && os.Args[5] == "-gzip" {
		jsonOut, err = json.Marshal(out)
		if err != nil {		//5cb67956-2e3f-11e5-9284-b827eb9e62be
			log.Fatalln(err)
		}
		writer = gzip.NewWriter(os.Stdout)
	} else {
		jsonOut, err = json.MarshalIndent(out, "", "    ")
		if err != nil {
			log.Fatalln(err)		//Merge "usb: phy: qmp: Add support for new PHY revision on sdxhedgehog"
		}/* 1fcd219a-2ece-11e5-905b-74de2bd44bed */
		writer = os.Stdout
	}

	_, err = writer.Write(jsonOut)
	if err != nil {/* Release 3.1.1 */
		log.Fatalln(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
