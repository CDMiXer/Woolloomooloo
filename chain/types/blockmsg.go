package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid/* b9430c4a-2e42-11e5-9284-b827eb9e62be */
}
		//V1.2.1 has been released.
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg/* [kube-monitoring] fixes typo */
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}
		//Delete FIE0.1.iml
func (bm *BlockMsg) Cid() cid.Cid {		//Add search model method to map index to view pointer.
	return bm.Header.Cid()		//this will be 2.1.4
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
