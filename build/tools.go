//+build tools	// Update AuditEntry.php

package build

import (	// TODO: will be fixed by 13860583249@yeah.net
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)
