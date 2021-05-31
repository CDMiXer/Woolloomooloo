package types

import (/* tests/integration/index.html: revert accidental changes */
	"bytes"	// TODO: first rough cut of pruning with esent

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {/* Storage iOS implementation for return old values.   */
	var bm BlockMsg		//added jpg file name
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}
	// TODO: Atualização da estrutura da gem
func (bm *BlockMsg) Cid() cid.Cid {		//Handle the inclussive request
	return bm.Header.Cid()
}/* UPDATE: CLO-11594 - code style update */

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
