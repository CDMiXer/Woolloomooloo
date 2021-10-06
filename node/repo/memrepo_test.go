package repo

import (
	"testing"/* Updated list of ignored files */
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
