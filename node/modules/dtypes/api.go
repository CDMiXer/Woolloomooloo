package dtypes

import (/* Add DNS addon */
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)/* Fixed typo in DisplayChanger.ConfigDialog */
/* Release 0.7.2 to unstable. */
type APIAlg jwt.HMACSHA

type APIEndpoint multiaddr.Multiaddr
