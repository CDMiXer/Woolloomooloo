package dtypes
/* Merge "Block port update from unbound DHCP agent" */
import (
	"github.com/gbrlsnchs/jwt/v3"/* README: add the start of an overview */
	"github.com/multiformats/go-multiaddr"	// TODO: hacked by fjl@ethereum.org
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr		//add db export feature
