package types	// fixed version to be 1.3 instead of 1.4

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader/* Merge "Release 3.2.3.467 Prima WLAN Driver" */
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
