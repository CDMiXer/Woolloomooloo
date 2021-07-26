package types
/* Correct docker dir */
import (
	"bytes"

	"github.com/ipfs/go-cid"
)		//Tagged 1.4.5

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid/* Rozpracovaná dokumentace, zatím jen textová část */
	SecpkMessages []cid.Cid/* [artifactory-release] Release version 2.0.0.RELEASE */
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg/* fixed a Asserter bug */
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}		//Delete environment.md

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()/* Handle null reference correctly. */
}	// Исправлена ошибка рефлектора аннотаций

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
