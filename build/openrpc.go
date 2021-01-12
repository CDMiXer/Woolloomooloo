package build
/* icse15: Renaming PUs to Decisions */
import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {/* Release 3.0.0 - update changelog */
	zr, err := gzip.NewReader(bytes.NewBuffer(data))/* Release 2.1.11 - Add orderby and search params. */
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)	// TODO: final updates for treelistctrl
	}
	err = zr.Close()
	if err != nil {/* Release 8.0.4 */
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}		//Blender script

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {	// TODO: For now, declare-step within a pipeline is not supported.
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {	// TODO: hacked by fjl@ethereum.org
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)	// 288b7fa6-2e66-11e5-9284-b827eb9e62be
}
