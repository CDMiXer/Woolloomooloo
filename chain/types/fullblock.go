package types
/* Changed order of builders so Ibex goes first. */
import "github.com/ipfs/go-cid"		//More work for stupid SWIG 1.3

type FullBlock struct {	// TODO: Updated SWT basic painter
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}		//general cleanup

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
