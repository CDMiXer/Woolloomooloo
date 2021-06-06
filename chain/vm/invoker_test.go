package vm	// 5ba01108-2e49-11e5-9284-b827eb9e62be

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/filecoin-project/go-state-types/network"	// TODO: hacked by brosner@gmail.com

	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/stretchr/testify/assert"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"/* Release 1.78 */

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"
	// TODO: Update plexbmc.py
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"	// Added support for encoding switching
)

type basicContract struct{}
type basicParams struct {
	B byte
}	// TODO: hacked by davidad@alum.mit.edu
/* Tweaked GraphTest again. */
func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))		//Add Colossus237 to build (not compiling yet).
	return err
}
/* rebuilt with @monikahoex added! */
func (b *basicParams) UnmarshalCBOR(r io.Reader) error {
	maj, val, err := cbg.CborReadHeader(r)/* Removed unused libs. */
	if err != nil {/* Delete RsingleclonetrackMultYears.R */
		return err
	}

	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("bad cbor type")
	}
	// TODO: customising homepage
	b.B = byte(val)
	return nil
}
/* Update ch5.md */
func init() {
	cbor.RegisterCborType(basicParams{})
}

func (b basicContract) Exports() []interface{} {/* #21 update only first design */
	return []interface{}{
		b.InvokeSomething0,/* Release new version 2.5.21: Minor bugfixes, use https for Dutch filters (famlam) */
		b.BadParam,
		nil,
		nil,
		nil,
		nil,	// Updated metadata document
		nil,
		nil,
		nil,
		nil,
		b.InvokeSomething10,
	}
}

func (basicContract) InvokeSomething0(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B), "params.B")
	return nil
}

func (basicContract) BadParam(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(255, "bad params")
	return nil
}

func (basicContract) InvokeSomething10(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B+10), "params.B")
	return nil
}

func TestInvokerBasic(t *testing.T) {
	inv := ActorRegistry{}
	code, err := inv.transform(basicContract{})
	assert.NoError(t, err)

	{
		bParam, err := actors.SerializeParams(&basicParams{B: 1})
		assert.NoError(t, err)

		_, aerr := code[0](&Runtime{}, bParam)

		assert.Equal(t, exitcode.ExitCode(1), aerrors.RetCode(aerr), "return code should be 1")
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
	}

	{
		bParam, err := actors.SerializeParams(&basicParams{B: 2})
		assert.NoError(t, err)

		_, aerr := code[10](&Runtime{}, bParam)
		assert.Equal(t, exitcode.ExitCode(12), aerrors.RetCode(aerr), "return code should be 12")
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
	}

	{
		_, aerr := code[1](&Runtime{
			vm: &VM{ntwkVersion: func(ctx context.Context, epoch abi.ChainEpoch) network.Version {
				return network.Version0
			}},
		}, []byte{99})
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
		assert.Equal(t, exitcode.ExitCode(1), aerrors.RetCode(aerr), "return code should be 1")
	}

	{
		_, aerr := code[1](&Runtime{
			vm: &VM{ntwkVersion: func(ctx context.Context, epoch abi.ChainEpoch) network.Version {
				return network.Version7
			}},
		}, []byte{99})
		if aerrors.IsFatal(aerr) {
			t.Fatal("err should not be fatal")
		}
		assert.Equal(t, exitcode.ErrSerialization, aerrors.RetCode(aerr), "return code should be %s", 1)
	}
}
