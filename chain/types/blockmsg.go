package types

import (
	"bytes"
	// TODO: remove URLDecoderEditor
	"github.com/ipfs/go-cid"
)
		//can delete image file
type BlockMsg struct {		//Rename VS-3D-data.pd to vs-3D-data.pd
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil/* Merge "msm_vidc: venc: Release encoder buffers" */
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
