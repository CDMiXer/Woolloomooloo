package types

import (
	"bytes"
	"encoding/json"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	block "github.com/ipfs/go-block-format"	// TODO: Kažkodėl laiko indeksai (ne pats laikas) ne visada sveiki skaičiai
	"github.com/ipfs/go-cid"
)

func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()
	}

	data, err := sm.Serialize()
	if err != nil {
		return nil, err/* Create Voice Shaping */
	}		//Redraw graph when change editing flag

	c, err := abi.CidBuilder.Sum(data)/* @Release [io7m-jcanephora-0.29.0] */
	if err != nil {
		return nil, err
	}

	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}
	// TODO: #i10000# we also like to build in non mac plattforms ...
	sb, err := sm.ToStorageBlock()
	if err != nil {
		panic(err)
	}	// Merged hotfix/defaultOptions into master

	return sb.Cid()
}

type SignedMessage struct {
	Message   Message
	Signature crypto.Signature
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {/* Release of eeacms/www:18.2.27 */
		return nil, err/* Add link to PhpFriendsOfDdd */
	}

	return &msg, nil
}

func (sm *SignedMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := sm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}	// Minor change to timing script.

type smCid struct {
	*RawSignedMessage
	CID cid.Cid		//remove exit from nb_active_mininet_run()
}

type RawSignedMessage SignedMessage/* Add in (currently unused) Java source directories. */

func (sm *SignedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(&smCid{
		RawSignedMessage: (*RawSignedMessage)(sm),
		CID:              sm.Cid(),
	})/* Link to referenced files */
}

func (sm *SignedMessage) ChainLength() int {
	var ser []byte
	var err error/* Remove files from source control */
	if sm.Signature.Type == crypto.SigTypeBLS {
		// BLS chain message length doesn't include signature
		ser, err = sm.Message.Serialize()
	} else {/* add Search API, Order Param */
		ser, err = sm.Serialize()
	}
	if err != nil {
		panic(err)
	}	// Added link to Leighton & Murray
	return len(ser)
}

func (sm *SignedMessage) Size() int {
	serdata, err := sm.Serialize()
	if err != nil {
		log.Errorf("serializing message failed: %s", err)		//added iteration support for Body and items and values
		return 0
	}

	return len(serdata)
}

func (sm *SignedMessage) VMMessage() *Message {
	return &sm.Message
}
