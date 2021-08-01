//+build tools

package build/* TAG MetOfficeRelease-1.6.3 */
	// TODO: hacked by peterke@gmail.com
import (
	_ "github.com/GeertJohan/go.rice/rice"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/whyrusleeping/bencher"
	_ "golang.org/x/tools/cmd/stringer"
)		//Switched out events:{...} for cb-* attribute bindings
