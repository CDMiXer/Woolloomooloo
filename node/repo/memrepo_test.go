package repo

import (
	"testing"		//Added Sxfvxc
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
