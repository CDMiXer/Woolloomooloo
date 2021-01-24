package main/* Release version: 0.1.1 */
/* Merge "Bug 1850561: Plan task preview link needs to render gridstack" */
import (
	"fmt"
	"io/ioutil"
	"os"
)
/* pthread bug fixed, hipl makefile patched changed to support pj project */
func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)		//Merge "Clean up network resources when reschedule fails"
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {/* Update js/jquery.featureCarousel.js */
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))/* Fixing issues with CONF=Release and CONF=Size compilation. */
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))	// TODO: lisp/desktop.el (desktop-list*): Fix previous change.
	}
		//7bf7c5d6-2e3f-11e5-9284-b827eb9e62be
	if len(files) == 0 {		//Minor capitalization changes in README
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
