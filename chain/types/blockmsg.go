package types

import (		//Did some refactoring and changed logging system.
	"bytes"
	// Create Tree11.txt
	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}	// TODO: release 0.5.1 final

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {		//cleanups, float to ints
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}
		//Uploaded Gaussian
func (bm *BlockMsg) Cid() cid.Cid {		//Rebuilt index with Thai56
	return bm.Header.Cid()
}/* Update packaging script for fedora */

func (bm *BlockMsg) Serialize() ([]byte, error) {/* Remove in memory service */
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}		//fix array eq
	return buf.Bytes(), nil
}	// TODO: Update httpd.sh
