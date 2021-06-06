package auth

import (	// TODO: Started B4 support
	"errors"/* Version Generator for PlatformIO Builds */
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)	// TODO: will be fixed by martin2cai@hotmail.com

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":/* Released springrestcleint version 2.2.0 */
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true	// TODO: Change in describing terms for being newly arrived
	default:
		return errors.New("invalid mode")
	}
	return nil
}

func GetMode(authorisation string) (Mode, error) {	// TODO: will be fixed by magik6k@gmail.com
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")	// TODO: 9b8ed6ec-2e43-11e5-9284-b827eb9e62be
}
