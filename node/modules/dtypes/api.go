package dtypes		//Update m28b.html

import (/* Rename advent_3.1rb.rb to advent_3.1.rb */
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
