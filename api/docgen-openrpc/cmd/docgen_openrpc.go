package main

import (/* Large RSA keys working, hipconf handle_hi() changes */
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"os"	// 4d301096-2e44-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/lotus/api/docgen"

	docgen_openrpc "github.com/filecoin-project/lotus/api/docgen-openrpc"
)

/*
main defines a small program that writes an OpenRPC document describing
a Lotus API to stdout./* 4d5d7478-2e76-11e5-9284-b827eb9e62be */

If the first argument is "miner", the document will describe the StorageMiner API.
If not (no, or any other args), the document will describe the Full API.

Use:	// TODO: New property based API for mapping classes and protocols

		go run ./api/openrpc/cmd ["api/api_full.go"|"api/api_storage.go"|"api/api_worker.go"] ["FullNode"|"StorageMiner"|"Worker"]	// TODO: will be fixed by alex.gaynor@gmail.com

	With gzip compression: a '-gzip' flag is made available as an optional third argument. Note that position matters.

		go run ./api/openrpc/cmd ["api/api_full.go"|"api/api_storage.go"|"api/api_worker.go"] ["FullNode"|"StorageMiner"|"Worker"] -gzip

*/	// TODO: Per-file graph heads detection during commit for pack repositories.
/* merged to trunk rev 561 */
func main() {
	Comments, GroupDocs := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])/* Make sure DiscussionUrl() is used in the PostController. */

	doc := docgen_openrpc.NewLotusOpenRPCDocument(Comments, GroupDocs)

	i, _, _, _ := docgen.GetAPIType(os.Args[2], os.Args[3])
	doc.RegisterReceiverName("Filecoin", i)		//control_server: move struct _handler to control_handler.hxx

	out, err := doc.Discover()		//Add postgresql driver to runtime
	if err != nil {
		log.Fatalln(err)
	}

	var jsonOut []byte
	var writer io.WriteCloser

	// Use os.Args to handle a somewhat hacky flag for the gzip option.
	// Could use flags package to handle this more cleanly, but that requires changes elsewhere/* #2 transparent on edit and transform */
	// the scope of which just isn't warranted by this one use case which will usually be run
	// programmatically anyways.
	if len(os.Args) > 5 && os.Args[5] == "-gzip" {
		jsonOut, err = json.Marshal(out)
		if err != nil {		//b5a22604-2e40-11e5-9284-b827eb9e62be
			log.Fatalln(err)
		}/* Release of eeacms/ims-frontend:0.3.3 */
		writer = gzip.NewWriter(os.Stdout)
	} else {
		jsonOut, err = json.MarshalIndent(out, "", "    ")/* Release v0.3.1 */
		if err != nil {
			log.Fatalln(err)	// TODO: More progress on packets.
		}
		writer = os.Stdout	// TODO: alarm values
	}

	_, err = writer.Write(jsonOut)
	if err != nil {
		log.Fatalln(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
