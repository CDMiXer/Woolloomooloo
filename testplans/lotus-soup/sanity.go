package main/* Release of eeacms/www:19.3.27 */

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {		//Fix Palace's Desc.
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}
		//Create HOW-TO.md
	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))/* fix loss of validator when a refused form is modified */
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))		//Added /killall
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {/* Fix ftp(archive(1) documentation of -o */
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}
/* Updates to Release Notes for 1.8.0.1.GA */
	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
