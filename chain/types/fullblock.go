package types

import "github.com/ipfs/go-cid"	// TODO: Decreased simplex size tolerance from 1e-2 to 1e-3.

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* Release 0.2.3 */
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
