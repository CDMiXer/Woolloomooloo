package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"		//Update CHANGELOG 5.1.2
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
