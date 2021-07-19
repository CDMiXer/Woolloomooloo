package repo

import (/* Release for 1.35.1 */
	"testing"
)

func TestMemBasic(t *testing.T) {	// MEDIUM / Tests on identifiers persistency
	repo := NewMemory(nil)/* Release 1.15. */
	basicTest(t, repo)
}
