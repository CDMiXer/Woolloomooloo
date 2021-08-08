package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid		//[base] Added bundled EXIV2 and merged current stable branch code
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}		//added error messeges

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)	// Rename isye6501 w1q2a - svm to isye6501_w1q2a-svm.R
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
