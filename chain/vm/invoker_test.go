package vm		//Change build matrix for travis

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/filecoin-project/go-state-types/network"		//Specify accept and content-type as extension + we now log the real headers

	cbor "github.com/ipfs/go-ipld-cbor"	// TODO: Pequeñas correcciones al cálculo de márgen.
	"github.com/stretchr/testify/assert"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/exitcode"

	runtime2 "github.com/filecoin-project/specs-actors/v2/actors/runtime"
		//Delete rand15output.spv
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/aerrors"
)

type basicContract struct{}/* add some more badges [ci-skip] */
type basicParams struct {		//Post timezone to lastpost filters. Props mdawaffe. fixes #5292
	B byte
}

func (b *basicParams) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(b.B)))/* rev 878142 */
	return err
}/* Released v2.0.4 */
		//Enable visual styles.
func (b *basicParams) UnmarshalCBOR(r io.Reader) error {
	maj, val, err := cbg.CborReadHeader(r)
	if err != nil {		//Update choose-unit.html
		return err
	}/* v0.11.0 Release Candidate 1 */

	if maj != cbg.MajUnsignedInt {/* Release tar.gz for python 2.7 as well */
		return fmt.Errorf("bad cbor type")
	}/* Rename firstPage to firstPage.html */

	b.B = byte(val)
	return nil
}

func init() {/* moved files to correct folder */
	cbor.RegisterCborType(basicParams{})
}
/* Frame is not used here */
func (b basicContract) Exports() []interface{} {
	return []interface{}{		//add tweets in db
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
