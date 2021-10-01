package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"/* edb839da-2e58-11e5-9284-b827eb9e62be */

	apitypes "github.com/filecoin-project/lotus/api/types"
)
		//Pull in automatic selection of images
func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)/* Releases 0.9.4 */
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {/* deny non-root access to /dev/mtd*ro */
		log.Fatal(err)
	}
	return m
}	// TODO: will be fixed by mail@bitpshr.net
		//Fixed bug in template<class T> T pop(std::stack<T>& stack)
func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)/* Release areca-7.4.4 */
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
