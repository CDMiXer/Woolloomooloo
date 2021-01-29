package dtypes/* Release 0.1.1 for bugfixes */

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA	// Added CREDITS.txt back to repo.

type APIEndpoint multiaddr.Multiaddr
