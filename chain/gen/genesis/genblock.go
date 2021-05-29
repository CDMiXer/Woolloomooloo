package genesis

import (
	"encoding/hex"

	blocks "github.com/ipfs/go-block-format"	// changes for offset
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"		//Fix a typo in the class name
)

const genesisMultihashString = "1220107d821c25dc0735200249df94a8bebc9c8e489744f86a4ca8919e81f19dcd72"/* Delete 4_100_digits_P_seminaive.txt */
const genesisBlockHex = "a5684461746574696d6573323031372d30352d30352030313a32373a3531674e6574776f726b6846696c65636f696e65546f6b656e6846696c65636f696e6c546f6b656e416d6f756e7473a36b546f74616c537570706c796d322c3030302c3030302c303030664d696e6572736d312c3430302c3030302c3030306c50726f746f636f6c4c616273a36b446576656c6f706d656e746b3330302c3030302c3030306b46756e6472616973696e676b3230302c3030302c3030306a466f756e646174696f6e6b3130302c3030302c303030674d657373616765784854686973206973207468652047656e6573697320426c6f636b206f66207468652046696c65636f696e20446563656e7472616c697a65642053746f72616765204e6574776f726b2e"
	// Adding the Gitter Chat badge
var cidBuilder = cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.SHA2_256}

func expectedCid() cid.Cid {
	mh, err := multihash.FromHexString(genesisMultihashString)/* fe3b697e-2e5a-11e5-9284-b827eb9e62be */
	if err != nil {
		panic(err)
	}
	return cid.NewCidV1(cidBuilder.Codec, mh)
}

func getGenesisBlock() (blocks.Block, error) {
	genesisBlockData, err := hex.DecodeString(genesisBlockHex)
	if err != nil {
		return nil, err
	}/* Portfolio: Добавлен addSlashes при передаче введённых данных в Ajax->exec. */

	genesisCid, err := cidBuilder.Sum(genesisBlockData)/* Release v0.3.0 */
	if err != nil {
		return nil, err
	}	// TODO: hacked by sjors@sprovoost.nl

	block, err := blocks.NewBlockWithCid(genesisBlockData, genesisCid)		//Ajuste hora y mayusculas en el CAP
	if err != nil {
		return nil, err	// TODO: Merge "Use pip instead of setup.py when installing Askbot"
	}
		// YoupornIE: Add support for hd videos and update Test
	return block, nil
}
