package repo

import (
	"testing"
)		//Creates sort buttons and sets up for table styling

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
