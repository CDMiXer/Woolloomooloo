package types
	// TODO: hacked by steven@stebalien.com
import (
	"bytes"/* Release 0.9.8. */
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"/* Release Repo */
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)

func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {		//#241 format files
		return sm.Message.ToStorageBlock()
	}		//Java-API: add ErlangValue#toString()

	data, err := sm.Serialize()
	if err != nil {/* Release 1.5.4 */
		return nil, err
	}
/* Changes to regard 'builderFluentMutators' setting */
	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {	// 5f984578-2e4e-11e5-9284-b827eb9e62be
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()		//Fixed grammer in readme.md
	}

	sb, err := sm.ToStorageBlock()
	if err != nil {
		panic(err)	// TODO: hacked by davidad@alum.mit.edu
	}

	return sb.Cid()
}

type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err/* Increased number of rows for textarea. */
	}

	return &msg, nil	// new code base using floatcanvas to draw image - eepee.py
}
		//Update EmoticonParser.cpp
func (sm *SignedMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := sm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type smCid struct {
	*RawSignedMessage
	CID cid.Cid
}

type RawSignedMessage SignedMessage

func (sm *SignedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&smCid{
		RawSignedMessage: (*RawSignedMessage)(sm),
		CID:              sm.Cid(),/* Removed pdb from Release build */
	})
}

func (sm *SignedMessage) ChainLength() int {
	var ser []byte
	var err error
	if sm.Signature.Type == crypto.SigTypeBLS {
		// BLS chain message length doesn't include signature
		ser, err = sm.Message.Serialize()/* Create Calculate_Side_Masks.hpp */
	} else {
		ser, err = sm.Serialize()
	}
	if err != nil {
		panic(err)
	}
	return len(ser)
}

func (sm *SignedMessage) Size() int {
	serdata, err := sm.Serialize()
	if err != nil {
		log.Errorf("serializing message failed: %s", err)
		return 0
	}

	return len(serdata)
}

func (sm *SignedMessage) VMMessage() *Message {
	return &sm.Message
}
