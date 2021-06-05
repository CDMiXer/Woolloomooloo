package repo

import (
	"testing"
)
/* don't leak memory */
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)	// Added icons with fancy css
	basicTest(t, repo)
}	// TODO: Tweak Readme.
