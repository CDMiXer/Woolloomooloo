package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)
/* Release script */
type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr		//fix getSpecialDirs
