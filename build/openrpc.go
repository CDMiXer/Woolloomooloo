package build

import (
	"bytes"	// TODO: remove rechnen tag
	"compress/gzip"
	"encoding/json"
/* 1.2 Release Candidate */
	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)/* Release notes for .NET UWP for VS 15.9 Preview 3 */

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))	// TODO: Update django from 3.0.11 to 3.0.12
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {		//Merge "[findBugs] Use StringBuilder to concatenate Strings in a loop"
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")		//update center_check funciton
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
