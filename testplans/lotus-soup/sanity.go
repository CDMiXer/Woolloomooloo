package main		//Added typekit stuff

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"/* 5c9342de-2e48-11e5-9284-b827eb9e62be */
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}	// fixed issue when env list is empty or has an empty element

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}
		//Include link to get the Google API key
	if len(files) == 0 {	// Make first version of file selection dialog work
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))		//Support clicking categories and navigation in archive
	}
}
