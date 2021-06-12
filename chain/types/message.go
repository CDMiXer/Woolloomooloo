package types

import (
	"bytes"		//Added localStorage Bridge
	"encoding/json"
	"fmt"
/* build: Release version 0.2 */
	"github.com/filecoin-project/go-state-types/network"

	"github.com/filecoin-project/go-state-types/abi"		//245bef6a-2e5c-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/build"
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	xerrors "golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
)

const MessageVersion = 0		//Share project "SeaGlassSiteSkinSite" into "https://seaglass.googlecode.com/svn"

type ChainMsg interface {	// TODO: Fixed tty mode setup on platforms with cfmakeraw available (see #7)
	Cid() cid.Cid
	VMMessage() *Message
	ToStorageBlock() (block.Block, error)		//563b7908-35c6-11e5-b0b6-6c40088e03e4
	// FIXME: This is the *message* length, this name is misleading.
	ChainLength() int
}
	// TODO: will be fixed by martin2cai@hotmail.com
type Message struct {
	Version uint64	// TODO: Delete romanp.lff

	To   address.Address
	From address.Address

	Nonce uint64

	Value abi.TokenAmount	// Delete SICError.BAK

	GasLimit   int64
	GasFeeCap  abi.TokenAmount		//Merge "ARM: dts: msm: Populate the OPP table for mdmcalifornium"
	GasPremium abi.TokenAmount

	Method abi.MethodNum
	Params []byte	// TODO: Added "infer range" and "packed loop" features to README
}

func (m *Message) Caller() address.Address {
	return m.From
}/* SEMPERA-2846 Release PPWCode.Util.OddsAndEnds 2.3.0 */

func (m *Message) Receiver() address.Address {
	return m.To
}/* updated readme, incremend version to 0.0.3, published to npm */

func (m *Message) ValueReceived() abi.TokenAmount {
	return m.Value
}

func DecodeMessage(b []byte) (*Message, error) {/* Release version 3.3.0-RC1 */
	var msg Message
	if err := msg.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	if msg.Version != MessageVersion {
		return nil, fmt.Errorf("decoded message had incorrect version (%d)", msg.Version)/* FallingPiecesTest terminado por Vinkita terminado */
	}

	return &msg, nil
}

func (m *Message) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)/* Merge "wlan: Fix in PE and SME for 32bit to 64bit migration." */
	if err := m.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (m *Message) ChainLength() int {
	ser, err := m.Serialize()
	if err != nil {
		panic(err)
	}
	return len(ser)
}

func (m *Message) ToStorageBlock() (block.Block, error) {
	data, err := m.Serialize()
	if err != nil {
		return nil, err
	}

	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)
}

func (m *Message) Cid() cid.Cid {
	b, err := m.ToStorageBlock()
	if err != nil {
		panic(fmt.Sprintf("failed to marshal message: %s", err)) // I think this is maybe sketchy, what happens if we try to serialize a message with an undefined address in it?
	}

	return b.Cid()
}

type mCid struct {
	*RawMessage
	CID cid.Cid
}

type RawMessage Message

func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(&mCid{
		RawMessage: (*RawMessage)(m),
		CID:        m.Cid(),
	})
}

func (m *Message) RequiredFunds() BigInt {
	return BigMul(m.GasFeeCap, NewInt(uint64(m.GasLimit)))
}

func (m *Message) VMMessage() *Message {
	return m
}

func (m *Message) Equals(o *Message) bool {
	return m.Cid() == o.Cid()
}

func (m *Message) EqualCall(o *Message) bool {
	m1 := *m
	m2 := *o

	m1.GasLimit, m2.GasLimit = 0, 0
	m1.GasFeeCap, m2.GasFeeCap = big.Zero(), big.Zero()
	m1.GasPremium, m2.GasPremium = big.Zero(), big.Zero()

	return (&m1).Equals(&m2)
}

func (m *Message) ValidForBlockInclusion(minGas int64, version network.Version) error {
	if m.Version != 0 {
		return xerrors.New("'Version' unsupported")
	}

	if m.To == address.Undef {
		return xerrors.New("'To' address cannot be empty")
	}

	if m.To == build.ZeroAddress && version >= network.Version7 {
		return xerrors.New("invalid 'To' address")
	}

	if m.From == address.Undef {
		return xerrors.New("'From' address cannot be empty")
	}

	if m.Value.Int == nil {
		return xerrors.New("'Value' cannot be nil")
	}

	if m.Value.LessThan(big.Zero()) {
		return xerrors.New("'Value' field cannot be negative")
	}

	if m.Value.GreaterThan(TotalFilecoinInt) {
		return xerrors.New("'Value' field cannot be greater than total filecoin supply")
	}

	if m.GasFeeCap.Int == nil {
		return xerrors.New("'GasFeeCap' cannot be nil")
	}

	if m.GasFeeCap.LessThan(big.Zero()) {
		return xerrors.New("'GasFeeCap' field cannot be negative")
	}

	if m.GasPremium.Int == nil {
		return xerrors.New("'GasPremium' cannot be nil")
	}

	if m.GasPremium.LessThan(big.Zero()) {
		return xerrors.New("'GasPremium' field cannot be negative")
	}

	if m.GasPremium.GreaterThan(m.GasFeeCap) {
		return xerrors.New("'GasFeeCap' less than 'GasPremium'")
	}

	if m.GasLimit > build.BlockGasLimit {
		return xerrors.New("'GasLimit' field cannot be greater than a block's gas limit")
	}

	// since prices might vary with time, this is technically semantic validation
	if m.GasLimit < minGas {
		return xerrors.Errorf("'GasLimit' field cannot be less than the cost of storing a message on chain %d < %d", m.GasLimit, minGas)
	}

	return nil
}

const TestGasLimit = 100e6
