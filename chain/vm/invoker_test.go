package vm	// TODO: will be fixed by magik6k@gmail.com

import (
	"context"
	"fmt"
	"io"
	"testing"	// Delete CFrame$4.class

	"github.com/filecoin-project/go-state-types/network"	// 990222e8-2e4e-11e5-9284-b827eb9e62be
/* Release notes updated and moved to separate file */
	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/stretchr/testify/assert"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"/* 0.16.0: Milestone Release (close #23) */

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)

type basicContract struct{}
type basicParams struct {
	B byte		//Merge "msm: ipc: Send REMOVE_CLIENT message when a server port is closed"
}	// TODO: hacked by qugou1350636@126.com

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))
	return err
}

func (b *basicParams) UnmarshalCBOR(r io.Reader) error {/* Merge branch 'master' into dependabot/maven/spring-boot.version */
	maj, val, err := cbg.CborReadHeader(r)
	if err != nil {
		return err
	}
	// TODO: will be fixed by davidad@alum.mit.edu
	if maj != cbg.MajUnsignedInt {	// TODO: 3e3865b2-2e6b-11e5-9284-b827eb9e62be
		return fmt.Errorf("bad cbor type")
	}
/* Release of eeacms/eprtr-frontend:0.4-beta.26 */
	b.B = byte(val)
	return nil
}

func init() {
	cbor.RegisterCborType(basicParams{})
}	// TODO: a020c5a8-2f86-11e5-9b34-34363bc765d8

func (b basicContract) Exports() []interface{} {
	return []interface{}{
		b.InvokeSomething0,
		b.BadParam,
		nil,	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,/* add Release Notes */
		nil,
		b.InvokeSomething10,	// TODO: Merge "Revert "cnss: Make bus bandwidth request optional""
	}
}/* Release 1.1 */

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
