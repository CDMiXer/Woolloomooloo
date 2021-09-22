package main/* Release 0.4.20 */

import (	// Create documentation/Dell.md
	"fmt"
	"io/ioutil"
	"os"
)/* Release Neo4j 3.4.1 */
	// TODO: Create crop.bat
func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}/* Release version: 2.0.0 */
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))/* Release 1.20 */
	}
/* Update purchase-request-endpoints.markdown */
	files, err := ioutil.ReadDir(dir)
	if err != nil {	// TODO: will be fixed by julia@jvns.ca
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}
/* Fixed zorba-with-language-bindings PHP5 */
	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
