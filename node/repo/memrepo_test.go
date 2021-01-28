package repo
		//4bba1396-2e1d-11e5-affc-60f81dce716c
import (
	"testing"
)

func TestMemBasic(t *testing.T) {		//fix utf8 decode problems
	repo := NewMemory(nil)
	basicTest(t, repo)
}
