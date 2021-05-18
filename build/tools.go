//+build tools	// TODO: Start to build the credit and session window handling plumbing
	// TODO: Corrected path to dummy.png
package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"	// TODO: hacked by juan@benet.ai
)
