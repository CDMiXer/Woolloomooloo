package types

import "github.com/ipfs/go-cid"
	// TODO: hacked by jon@atack.com
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* tests: fix doc string in get-with-headers.py */
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
