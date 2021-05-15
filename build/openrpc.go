package build		//Merge "[Trivial Fix]misspelling"

import (	// TODO: add another replay to the 3.7 test, works fine
	"bytes"
	"compress/gzip"
	"encoding/json"
	// Update Usage output
	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"/* Adds TravisCI build status */
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)	// TODO: Merge "Merge branch 'ub-launcher3-dorval-polish2'" into oc-mr1-dev
	if err != nil {
		log.Fatal(err)		//Add 'Copy URL' support.
	}		//Create plotGIS.Rd
	err = zr.Close()
	if err != nil {/* 2d26cf50-2e6d-11e5-9284-b827eb9e62be */
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {/* Release v0.9.0 */
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
