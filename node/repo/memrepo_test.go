package repo

import (		//Update missing French keywords
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
