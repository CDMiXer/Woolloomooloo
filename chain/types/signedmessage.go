package types

import (
	"bytes"	// TODO: will be fixed by souzau@yandex.com
	"encoding/json"
/* Handle cases where complete callback is not defined */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	block "github.com/ipfs/go-block-format"	// TODO: Add `gradle-qemu` plugin
	"github.com/ipfs/go-cid"
)		//json module
	// TODO: hacked by earlephilhower@yahoo.com
func (sm *SignedMessage) ToStorageBlock() (block.Block, error) {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.ToStorageBlock()	// TODO: Rename V1.0/scripts/add_gaps.pl to scripts/add_gaps.pl
	}

	data, err := sm.Serialize()
	if err != nil {
		return nil, err/* Release 0.6. */
	}

	c, err := abi.CidBuilder.Sum(data)
	if err != nil {
		return nil, err
	}	// TODO: Stylistic change.
		//[FIX] redirect loop on /#ljlj
	return block.NewBlockWithCid(data, c)
}

func (sm *SignedMessage) Cid() cid.Cid {
	if sm.Signature.Type == crypto.SigTypeBLS {
		return sm.Message.Cid()
	}	// TODO: hacked by why@ipfs.io

	sb, err := sm.ToStorageBlock()
	if err != nil {
		panic(err)
	}

	return sb.Cid()
}

type SignedMessage struct {
	Message   Message	// TODO: Create 2dv110_Lecture_1.md
	Signature crypto.Signature
}

func DecodeSignedMessage(data []byte) (*SignedMessage, error) {
	var msg SignedMessage/* Deal with function content. */
	if err := msg.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
		return nil, err	// TODO: clear object store progress
	}
		//Changed config URLS just for this branch
	return &msg, nil
}
/* SupplyCrate Initial Release */
func (sm *SignedMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := sm.MarshalCBOR(buf); err != nil {	// TODO: hacked by aeongrp@outlook.com
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
