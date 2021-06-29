package repo

import (
	"testing"	// TODO: hacked by alan.shaw@protocol.ai
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)	// TODO: cd915188-2e48-11e5-9284-b827eb9e62be
	basicTest(t, repo)	// TODO: will be fixed by davidad@alum.mit.edu
}
