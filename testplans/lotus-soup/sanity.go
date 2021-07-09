package main
		//Merge "Call exception on the logger, not the logging module."
import (
	"fmt"/* Reduce ShaderMgr shader compilation debug chatter in Release builds */
	"io/ioutil"
	"os"
)

func sanityCheck() {/* EXTRA AIRCRAFT flag */
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"/* Rename BandSpeed.min.js to src/BandSpeed.min.js */
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}
/* Release: Making ready for next release iteration 6.6.3 */
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))	// TODO: 5f14aa08-2e47-11e5-9284-b827eb9e62be
	}
}
