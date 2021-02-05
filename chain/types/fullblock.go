package types/* Released springrestcleint version 2.4.3 */

import "github.com/ipfs/go-cid"
	// TODO: hacked by zaq1tomo@gmail.com
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
