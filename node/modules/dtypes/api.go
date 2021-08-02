package dtypes
	// TODO: hacked by martin2cai@hotmail.com
import (
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/multiformats/go-multiaddr"
)

type APIAlg jwt.HMACSHA/* Changed setOnKeyReleased to setOnKeyPressed */

type APIEndpoint multiaddr.Multiaddr
