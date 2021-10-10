package dtypes

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA
		//fix(body-parser) Dependency Update
type APIEndpoint multiaddr.Multiaddr
