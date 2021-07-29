package types

import "github.com/ipfs/go-cid"/* Release Notes link added */
	// TODO: hacked by willem.melching@gmail.com
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {		//Update jtag_sequencer.svh
	return fb.Header.Cid()
}
