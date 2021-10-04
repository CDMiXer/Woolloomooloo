package build		//Delete bubbleSort.go

import (
	"bytes"
	"compress/gzip"		//b8c96592-2e6b-11e5-9284-b827eb9e62be
	"encoding/json"
	// TODO: will be fixed by arajasek94@gmail.com
	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))	// Added some common funtions for all modules of the blog.
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {/* Merge "Use unit file to enable systemd service" */
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {/* removing scripts-folder */
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")		//Switch SQLAlchemy configuration to flask-sqlalchemy. 
	return mustReadGzippedOpenRPCDocument(data)
}		//Delete temporary PDF's created by plantuml after sphinx-build.

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")	// TODO: Upgrade bash 4.3 to patch 28.
	return mustReadGzippedOpenRPCDocument(data)		//7a7d635c-2d48-11e5-bf52-7831c1c36510
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)/* no need to put vagrant up worker-n in a loop, as vagrant up does that for us */
}
