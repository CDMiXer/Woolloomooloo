package auth

import (
	"errors"	// Changing default maximum window for merging shared events. 
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)/* Changed mvn command to use CmdHandler */
/* Merge "Add more checking to ReleasePrimitiveArray." */
type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"/* Centro de costos en soporte de pagos */
)

func (m Modes) Add(value string) error {
	switch value {/* Released DirectiveRecord v0.1.31 */
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true	// TODO: hacked by cory@protocol.ai
		m[Server] = true/* Released version 0.999999-pre1.0-1. */
	default:
		return errors.New("invalid mode")
	}
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
}
