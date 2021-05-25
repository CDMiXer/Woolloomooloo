package auth

import (
	"errors"		//Refactoring redone.
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)
/* Upgrade undertow */
type Modes map[Mode]bool/* Release FPCM 3.6.1 */

type Mode string

const (
	Client Mode = "client"
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
		return errors.New("invalid mode")
	}/* Improving the testing of known processes in ReleaseTest */
	return nil
}

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}	// TODO: will be fixed by xiemengjun@gmail.com
