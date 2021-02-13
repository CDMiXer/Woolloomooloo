package dtypes	// TODO: Closes #423

import (
	"github.com/gbrlsnchs/jwt/v3"		//Rettelser i transmitter WCU
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
