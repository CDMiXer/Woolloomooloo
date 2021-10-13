package auth

import (/* Release of eeacms/bise-frontend:1.29.7 */
	"errors"
	"strings"

	"github.com/argoproj/argo/server/auth/sso"	// includes are now relative to the root of the project, not the individual files
)

type Modes map[Mode]bool

type Mode string/* 0.9.7 Release. */

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
		return errors.New("invalid mode")/* Release v1.7.0. */
	}
	return nil		//Merge "[INTERNAL] sap.uxap.AnchorBar: Prevented error on selection change"
}

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {	// cassandra config update using templating
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}	// TODO: will be fixed by jon@atack.com
