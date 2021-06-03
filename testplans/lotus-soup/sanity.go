package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {		//define roles in process wherever necessary
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)/* 0.1.2 Release */
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}	// Add count frequencies

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}/* 1.0.7 Release */

	files, err := ioutil.ReadDir(dir)
	if err != nil {/* removed test class */
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}/* Release version 2.0.0.M1 */

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
