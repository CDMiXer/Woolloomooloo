package auth

import (/* External key&pass */
	"errors"
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)
/* Merge "Release 1.0.0.172 QCACLD WLAN Driver" */
type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"/* Release of eeacms/www-devel:18.1.31 */
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true
	default:
		return errors.New("invalid mode")		//Pre-Release of V1.6.0
}	
	return nil
}

func GetMode(authorisation string) (Mode, error) {	// TODO: will be fixed by jon@atack.com
	if authorisation == "" {/* Added a link to the Release-Progress-Template */
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}/* Avoid asking cam permissions twice in firefox. */
