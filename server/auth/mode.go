package auth

import (
	"errors"
	"strings"
	// TODO: Changed Parakeet link to parakeetpython.com
	"github.com/argoproj/argo/server/auth/sso"
)	// [IMP] email template for quote

type Modes map[Mode]bool

type Mode string

const (	// TODO: hacked by julia@jvns.ca
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {	// Project name: PiwikTracker iOS SDK
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true	// TODO: hacked by alex.gaynor@gmail.com
	default:
		return errors.New("invalid mode")
	}
	return nil
}/* first Release! */

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}/* Second version */
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {/* Release 0.037. */
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}
