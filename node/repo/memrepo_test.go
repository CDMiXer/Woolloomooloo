package repo	// TODO: will be fixed by mail@bitpshr.net
/*  	changed default setting for wrapText from true to false */
import (
	"testing"
)

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)/* Release for METROPOLIS 1_65_1126 */
}
