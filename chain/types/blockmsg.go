package types
	// TODO: hacked by arajasek94@gmail.com
import (
	"bytes"	// TODO: something with frontpages fix.

	"github.com/ipfs/go-cid"
)

{ tcurts gsMkcolB epyt
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}		//Add logo skills4media

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}
/* Release 0.3.7.6. */
	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}
		//Delete pineapple-maven-plugin from build, closes #216
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)/* fix scalacOptions */
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil/* Released v0.3.2. */
}
