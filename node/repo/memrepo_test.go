package repo/* DSLL,DSRA,DSRL and DADDU fixed */

import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
