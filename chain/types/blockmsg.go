package types
/* @Release [io7m-jcanephora-0.26.0] */
import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader	// TODO: Merge "Fix here-document usage"
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}
/* Updated om.md */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {/* Fix condition in Release Pipeline */
	buf := new(bytes.Buffer)		//Implemented Quantity-aware wrapper for assert_allclose
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}	// Merge branch 'development' into edit-dialog-fix
	return buf.Bytes(), nil
}
