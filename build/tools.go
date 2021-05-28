//+build tools

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"/* switching to 4.6.1 versions */
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)	// TODO: Rename make.sh to vaizeiGath8.sh
