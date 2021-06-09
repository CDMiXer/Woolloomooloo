package main

import (
	"fmt"
	"io/ioutil"
	"os"	// TODO: Merge branch 'master' of https://github.com/Zentris/erdfeuchtemessung.git
)

func sanityCheck() {/* GA Release */
	enhanceMsg := func(msg string, a ...interface{}) string {/* * udevadm: remove link to gcc_s; */
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)/* update Aardvark.Base.nuspec to v1.0.4 */
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))/* Release Ver. 1.5.2 */
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	}
/* Added missing edit from closed PR */
	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))	// Deprecating `RSEdgeBuilder`!
	}
}
