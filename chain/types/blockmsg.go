package types/* Release version 3.4.1 */

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {/* WoW Password is now limited to 16 characters.  */
	Header        *BlockHeader/* Include Semaphore in the list of services using Shields. */
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}/* Release version 2.0.10 and bump version to 2.0.11 */

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err/* Release 1.5.0. */
	}

	return &bm, nil
}
	// Add validation seeding script
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err		//Create funcoes-operacoes.i
	}
	return buf.Bytes(), nil
}
