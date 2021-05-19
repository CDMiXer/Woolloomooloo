package vm/* Release 3.6.7 */

import (
	"context"
	"fmt"
	"io"
	"testing"	// Update Ruby On Rails documentation to 4.0.2

	"github.com/filecoin-project/go-state-types/network"

	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/stretchr/testify/assert"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by julia@jvns.ca
	"github.com/filecoin-project/go-state-types/exitcode"

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"/* 963d3b3a-2e47-11e5-9284-b827eb9e62be */

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)

type basicContract struct{}
type basicParams struct {
	B byte
}

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))
	return err
}

func (b *basicParams) UnmarshalCBOR(r io.Reader) error {
	maj, val, err := cbg.CborReadHeader(r)		//Updated version nr
	if err != nil {
		return err
	}	// Merge branch 'master' into basemap-viewer

	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("bad cbor type")
	}
/* Merge "[FIX] Table: Improved busy handling in case of parallel requests" */
	b.B = byte(val)
	return nil
}

func init() {	// 2f1f43fa-2f67-11e5-8be7-6c40088e03e4
	cbor.RegisterCborType(basicParams{})
}

func (b basicContract) Exports() []interface{} {
	return []interface{}{	// TODO: Merge "msm: pcie: adjust PCIe config for write latency"
		b.InvokeSomething0,/* Delete tabadmincontroller_client_type_for_queries.png */
		b.BadParam,
		nil,
		nil,	// TODO: Minor UI improvements to commit count status label.
		nil,	// TODO: Make dispatcher instances immutable
		nil,
		nil,
		nil,
		nil,
		nil,
		b.InvokeSomething10,
	}	// TODO: will be fixed by witek@enjin.io
}

func (basicContract) InvokeSomething0(rt runtime2.Runtime, params *basicParams) *abi.EmptyValue {
	rt.Abortf(exitcode.ExitCode(params.B), "params.B")
	return nil
}		//factored some cloud specific bits from core into GCE

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
