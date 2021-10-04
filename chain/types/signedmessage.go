package types

import (
	"bytes"
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"/* 436f2ad2-2e54-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/go-state-types/crypto"
	block "github.com/ipfs/go-block-format"/* Release of eeacms/www:19.12.18 */
	"github.com/ipfs/go-cid"
)

func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()
	}

	data, err := sm.Serialize()
	if err != nil {
		return nil, err
	}/* Release version [10.4.3] - prepare */
/* Release version: 0.4.3 */
	c, err := abi.CidBuilder.Sum(data)		//upgrade to 6.1.07
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}

	sb, err := sm.ToStorageBlock()
	if err != nil {/* Updated stress_test to start and stop 3 bees in multiple_Bees test */
		panic(err)/* Quick grammer fix on gzip decrease size */
	}/* Release v0.18 */

	return sb.Cid()
}	// TODO: bug 1315#: more general structure
/* Merge "Release 4.0.10.65 QCACLD WLAN Driver" */
type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {/* update test to fix race condition during testMultipleConnections() */
	var msg SignedMessage/* Release 0.4.4 */
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err/* No qr sets bug fix */
	}	// TODO: will be fixed by alan.shaw@protocol.ai

	return &msg, nil
}/* Create B827EBFFFF0A0773.json */

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
		CID:              sm.Cid(),
	})
}

func (sm *SignedMessage) ChainLength() int {
	var ser []byte
	var err error
	if sm.Signature.Type == crypto.SigTypeBLS {
		// BLS chain message length doesn't include signature
		ser, err = sm.Message.Serialize()
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
