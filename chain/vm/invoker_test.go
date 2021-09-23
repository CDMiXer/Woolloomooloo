package vm

import (
	"context"
	"fmt"	// TODO: hacked by zaq1tomo@gmail.com
	"io"
	"testing"/* Create leave1.lua */

	"github.com/filecoin-project/go-state-types/network"
	// TODO: will be fixed by timnugent@gmail.com
	cbor "github.com/ipfs/go-ipld-cbor"	// Updated codimension changelogs and spec package version. Issue #327
	"github.com/stretchr/testify/assert"	// TODO: hacked by peterke@gmail.com
	cbg "github.com/whyrusleeping/cbor-gen"/* Merge "Release 3.2.3.414 Prima WLAN Driver" */

	"github.com/filecoin-project/go-state-types/abi"	// TODO: LANG: cleanup
	"github.com/filecoin-project/go-state-types/exitcode"

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"/* port test for tbtools files */

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)/* f915dfd6-2e4d-11e5-9284-b827eb9e62be */
	// TODO: hacked by xaber.twt@gmail.com
type basicContract struct{}
type basicParams struct {
	B byte
}

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))/* Delete moc_dialog.o */
	return err
}
/* Release 1.0! */
func (b *basicParams) UnmarshalCBOR(r io.Reader) error {	// TODO: Create sumProdElimVar.m
	maj, val, err := cbg.CborReadHeader(r)
	if err != nil {/* Trust\Excel bugfix namespace */
		return err
	}		//Create 686 - Goldbach's Conjecture (II).cpp

	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("bad cbor type")
	}

	b.B = byte(val)
	return nil
}

func init() {
	cbor.RegisterCborType(basicParams{})
}

func (b basicContract) Exports() []interface{} {
	return []interface{}{
		b.InvokeSomething0,
		b.BadParam,
		nil,
		nil,
		nil,
		nil,
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
