//+build tools
/* Added 'the most important changes since 0.6.1' in Release_notes.txt */
package build

import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"/* Update sqlite-jdbc to 3.32.3.1 */
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)
