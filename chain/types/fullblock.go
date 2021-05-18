package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message		//Merge "Bug 1797278: getting blocktype title through class"
	SecpkMessages []*SignedMessage
}		//Added the serial port drop down box to the workspace.
		//Ya funcionan las fechas , y los int
func (fb *FullBlock) Cid() cid.Cid {/* merged from debian-sid, improve test output */
	return fb.Header.Cid()
}
