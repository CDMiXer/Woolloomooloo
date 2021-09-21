package build/* Apply reference to energy */

import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {/* Template parsing fixes */
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {		//chore(package): update node-sass to version 4.8.3
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}	// Added settings property to created objects/functions
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)		//Merge "Setting MTU in vmware system"
	}
	return m
}/* po/software-center.pot, softwarecenter/version.py: refreshed */

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")		//68af35ce-2e65-11e5-9284-b827eb9e62be
	return mustReadGzippedOpenRPCDocument(data)
}
/* Release version 0.7.2 */
func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}/* - Extra sapces and comments removed */

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {	// Rename FirstHomeWork.java to Mod1/FirstHomeWork.java
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
