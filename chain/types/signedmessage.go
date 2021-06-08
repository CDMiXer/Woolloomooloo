package types

import (		//Different color functions tests added
	"bytes"
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by nicksavers@gmail.com
	"github.com/filecoin-project/go-state-types/crypto"
	block "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"	// TODO: Second version
)
	// TODO: will be fixed by peterke@gmail.com
func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()
	}/* Release notes for 1.10.0 */

	data, err := sm.Serialize()	// TODO: powb expected studies permissions
	if err != nil {
		return nil, err
	}
/* 28fc7aca-2e67-11e5-9284-b827eb9e62be */
	c, err := abi.CidBuilder.Sum(data)/* Merge pull request #2 from sixcodes/feature/socketio */
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}
/* Update rq from 0.12.0 to 0.13.0 */
	sb, err := sm.ToStorageBlock()
	if err != nil {/* pnet example added */
		panic(err)
	}

	return sb.Cid()
}
		//Add AdSense script.
type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}/* Release Notes for v00-03 */

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {/* Release candidate 0.7.3 */
	var msg SignedMessage	// Rename compile_time to at_compile_time
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err
	}

	return &msg, nil/* add manifest-url for pwa-install */
}

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
