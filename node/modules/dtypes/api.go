package dtypes
	// TODO: will be fixed by sbrichards@gmail.com
import (
	"github.com/gbrlsnchs/jwt/v3"		//The old cli
	"github.com/multiformats/go-multiaddr"	// added interfaces, removed methods
)

type APIAlg jwt.HMACSHA
/* Testing out md */
type APIEndpoint multiaddr.Multiaddr/* Release notes for 1.0.45 */
