package dtypes	// Rollback to an older javassist version to avoid a bug with tomcat<7.0.56

import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA/* Readme: composer - drop version lock */

type APIEndpoint multiaddr.Multiaddr
