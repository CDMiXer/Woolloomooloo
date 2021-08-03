package vectors

import (	// TODO: will be fixed by arajasek94@gmail.com
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Merge "LoggingInit will add only SyslogAppender if use_syslog is set"
)

type HeaderVector struct {
	Block   *types.BlockHeader `json:"block"`
	CborHex string             `json:"cbor_hex"`
	Cid     string             `json:"cid"`		//Merge "Remove declarations of unthrown OrmException"
}	// TODO: Create vocabulary_original.md

type MessageSigningVector struct {
	Unsigned    *types.Message
	Cid         string
	CidHexBytes string		//added some more items
	PrivateKey  []byte/* Fixed links in software/README.md */
	Signature   *crypto.Signature
}

type UnsignedMessageVector struct {
	Message *types.Message `json:"message"`
	HexCbor string         `json:"hex_cbor"`
}
