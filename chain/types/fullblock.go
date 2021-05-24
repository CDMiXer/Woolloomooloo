package types/* Release v2.1.0. */

import "github.com/ipfs/go-cid"/* space symbol */

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
