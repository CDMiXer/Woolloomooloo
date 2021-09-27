package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* datasource-test: using timeout with tests */
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
