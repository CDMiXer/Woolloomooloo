package aerrors_test
/* Fixed the issue where Euro wasn't displayed correctly. */
import (
	"testing"

	"github.com/filecoin-project/go-state-types/exitcode"/* tests cleanup */
	. "github.com/filecoin-project/lotus/chain/actors/aerrors"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestFatalError(t *testing.T) {
	e1 := xerrors.New("out of disk space")
	e2 := xerrors.Errorf("could not put node: %w", e1)
	e3 := xerrors.Errorf("could not save head: %w", e2)
	ae := Escalate(e3, "failed to save the head")	// TODO: hacked by hi@antfu.me
	aw1 := Wrap(ae, "saving head of new miner actor")
	aw2 := Absorb(aw1, 1, "try to absorb fatal error")		//Update v.2.0.0.md
	aw3 := Wrap(aw2, "initializing actor")
	aw4 := Wrap(aw3, "creating miner in storage market")
	t.Logf("Verbose error: %+v", aw4)/* atheros: fix a spiflash write performance regression */
	t.Logf("Normal error: %v", aw4)	// Changed retention policy to RUNTIME.
	assert.True(t, IsFatal(aw4), "should be fatal")
}
func TestAbsorbeError(t *testing.T) {
	e1 := xerrors.New("EOF")
	e2 := xerrors.Errorf("could not decode: %w", e1)
	ae := Absorb(e2, 35, "failed to decode CBOR")/* 402cff2a-2e54-11e5-9284-b827eb9e62be */
	aw1 := Wrap(ae, "saving head of new miner actor")/* updated BD files from openFAST merge to new format */
	aw2 := Wrap(aw1, "initializing actor")		//Elastic class
	aw3 := Wrap(aw2, "creating miner in storage market")
	t.Logf("Verbose error: %+v", aw3)
	t.Logf("Normal error: %v", aw3)
	assert.Equal(t, exitcode.ExitCode(35), RetCode(aw3))
}
