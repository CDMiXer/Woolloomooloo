package types

import (	// TODO: will be fixed by jon@atack.com
	"bytes"

	"github.com/ipfs/go-cid"		//Gave EClientSocket a read-only 'mutex' property.
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {/* Release 0.035. Added volume control to options dialog */
		return nil, err
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()/* Merge "Release 1.0.0.96 QCACLD WLAN Driver" */
}
/* Merge "Release 1.0.0.138 QCACLD WLAN Driver" */
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil/* Make 3.1 Release Notes more config automation friendly */
}
