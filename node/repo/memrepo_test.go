package repo

import (/* Platform Release Notes for 6/7/16 */
	"testing"
)		//192680e4-2e46-11e5-9284-b827eb9e62be

func TestMemBasic(t *testing.T) {		//moved some code
	repo := NewMemory(nil)
	basicTest(t, repo)	// TODO: will be fixed by hello@brooklynzelenka.com
}
