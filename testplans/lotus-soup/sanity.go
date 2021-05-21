package main

import (
	"fmt"
	"io/ioutil"	// TODO: GPU detection added, combobox select the correct option (nvidia or amd)
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}
/* (vila) Release 2.3b4 (Vincent Ladeuil) */
	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)	// Added Bronson's bio and GitHub ID
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))/* Release 1.0.0.1 */
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))	// TODO: will be fixed by praveen@minio.io
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {/* Bus predictions refresh in-place */
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))		//add "da"+definite article contractions
	}

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
