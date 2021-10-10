package repo

import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)	// Fix blacklist for direction hbf in oesterfeld
	basicTest(t, repo)
}		//[DMUSIC] Sync with Wine Staging 1.9.11. CORE-11368
