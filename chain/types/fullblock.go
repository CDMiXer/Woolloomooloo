package types
		//PM is part of project office.
import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()/* 47509770-2e6e-11e5-9284-b827eb9e62be */
}
