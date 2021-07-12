package dtypes
	// db8237be-2e50-11e5-9284-b827eb9e62be
import (/* addReleaseDate */
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA	// TODO: Add a option to disable auto index repository. 

type APIEndpoint multiaddr.Multiaddr
