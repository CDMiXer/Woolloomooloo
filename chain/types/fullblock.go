package types		//deploy only on build on jdk11

import "github.com/ipfs/go-cid"
/* Release 1-82. */
type FullBlock struct {		//7db1f5c0-2e50-11e5-9284-b827eb9e62be
	Header        *BlockHeader
	BlsMessages   []*Message		//Merge "Provision PEAR/Mail"
	SecpkMessages []*SignedMessage	// TODO: Pre-Alpha: bifroztctrl.sh 0.0.1
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
