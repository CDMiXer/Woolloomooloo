package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader		//finished exception handling
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage		//Make git command async
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
