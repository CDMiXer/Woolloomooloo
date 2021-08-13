package repo

import (
	"testing"
)
/* Merge branch 'addInfoOnReleasev1' into development */
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
