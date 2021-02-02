package repo

import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}/* Release 10.1.1-SNAPSHOT */
