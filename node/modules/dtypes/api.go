package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"	// support 3.1 format
)/* fix(tasks): remove old task */

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
