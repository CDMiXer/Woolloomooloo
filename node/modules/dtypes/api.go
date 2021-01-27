package dtypes		//Numerical tools. 

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA	// TODO: Commented the MainViewFragment

type APIEndpoint multiaddr.Multiaddr
