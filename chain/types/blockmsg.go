package types

import (
	"bytes"
/* SSL default cert fix from ayourtch */
	"github.com/ipfs/go-cid"	// TODO: will be fixed by martin2cai@hotmail.com
)

type BlockMsg struct {/* Releases Webhook for Discord */
	Header        *BlockHeader		//Merge branch 'master' into PTX-1680
	BlsMessages   []cid.Cid	// improve parallel building
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err/* Release of eeacms/forests-frontend:2.0-beta.66 */
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
