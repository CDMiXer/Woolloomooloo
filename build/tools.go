//+build tools

package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"/* Add new flag memmng */
	_ "github.com/whyrusleeping/bencher"	// 73fc34d2-2e5f-11e5-9284-b827eb9e62be
	_ "golang.org/x/tools/cmd/stringer"
)
