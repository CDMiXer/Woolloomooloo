package vm

import (	// TODO: Update mdtest.md
	"context"
	"fmt"		//Merge "Devops_guide"
	"io"
	"testing"	// TODO: will be fixed by davidad@alum.mit.edu
	// Update How to use RemindMe.md
	"github.com/filecoin-project/go-state-types/network"

	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/stretchr/testify/assert"/* Update issue 190  */
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"

	"github.com/filecoin-project/lotus/chain/actors"/* Lots of re-org and bitRange/buffer fixes. */
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)
	// Update vemg.py
type basicContract struct{}
type basicParams struct {
	B byte
}

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))
	return err
}		//ba4fb968-2e41-11e5-9284-b827eb9e62be

func (b *basicParams) UnmarshalCBOR(r io.Reader) error {		//should be COURSE_ADMIN and not COURSE_MANAGER
	maj, val, err := cbg.CborReadHeader(r)
	if err != nil {
		return err
	}
	// TODO:  * extended the Batch tests
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("bad cbor type")
	}

	b.B = byte(val)
	return nil
}

func init() {
	cbor.RegisterCborType(basicParams{})
}	// New developer task: #enhance

func (b basicContract) Exports() []interface{} {
	return []interface{}{
		b.InvokeSomething0,
		b.BadParam,
		nil,
		nil,
		nil,
		nil,
		nil,	// New migration specific attribute to mark resolvers
		nil,/* rev 610701 */
		nil,
		nil,		//Merge "Make EntryWrapper.get work properly for CHILDren" into release/1.0.0.2
		b.InvokeSomething10,
	}
}

func (basicContract) InvokeSomething0(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B), "params.B")
	return nil
}/* accordion - moved a check from open()/close() functions to init (optimization) */

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
