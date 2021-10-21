package dtypes

import (	// TODO: will be fixed by steven@stebalien.com
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
