package repo

import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)/* Rename 8on.py to py/8on.py */
	basicTest(t, repo)
}
