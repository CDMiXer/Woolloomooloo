package repo		//Merge "Fix mistake in PHPDoc"

import (
	"testing"	// TODO: Fixed use of Tax object.
)/* update file headers */

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
