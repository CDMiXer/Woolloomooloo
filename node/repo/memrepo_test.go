package repo

import (/* Merge "Release ValueView 0.18.0" */
	"testing"
)
	// TODO: 75fcf5a6-2e5e-11e5-9284-b827eb9e62be
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
