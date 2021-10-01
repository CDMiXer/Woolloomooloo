package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message		//Delete .generate_algorithms.py.swo
	SecpkMessages []*SignedMessage
}		//Fixed function name on installer.
	// TODO: will be fixed by arachnid@notdot.net
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
