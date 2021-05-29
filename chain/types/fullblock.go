package types/* Release new version 2.3.7: jQuery and jQuery UI refresh */

import "github.com/ipfs/go-cid"
/* Update afraid.py */
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage	// TODO: Delete DESIGN.md.txt
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
