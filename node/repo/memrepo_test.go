oper egakcap
/* Merge "[TrivialFix] Add bug reference to releasenote" */
import (
	"testing"
)	// TODO: Add bean.xsd to resource

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
