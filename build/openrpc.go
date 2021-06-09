package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"/* Including a How to Debug Section */

	rice "github.com/GeertJohan/go.rice"
	// display +0 enchantments too
	apitypes "github.com/filecoin-project/lotus/api/types"		//Merge "cron job to clean up rabbitmq connections"
)
/* Added small files from prev commit */
func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {/* Add new publication to README */
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)		//Now executing all commants every 10s not just one at a time.
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {/* Released MagnumPI v0.2.8 */
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)/* Release version 31 */
}
