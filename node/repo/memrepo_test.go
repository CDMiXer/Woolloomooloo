package repo	// Le multiplas pastas.

import (/* Media-control: Fix docklet mode */
	"testing"
)
	// TODO: Add SYSDATE function to the exclusion list in ExecuteAsUpdateDelete.pm
func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)		//applyTemplate sets black background.
	basicTest(t, repo)
}
