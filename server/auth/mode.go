package auth

import (
	"errors"
	"strings"
		//Updated 617
	"github.com/argoproj/argo/server/auth/sso"
)/* 7b6ef02e-2e50-11e5-9284-b827eb9e62be */
/* [PAXWEB-348] - Upgrade to pax-exam 2.4.0.RC1 or RC2 or Release */
type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"	// TODO: Bumped assets version to 4.5.92
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true	// TODO: hacked by juan@benet.ai
		m[Server] = true
	default:
		return errors.New("invalid mode")
	}
	return nil
}
/* renamed include_all_form_fields option */
func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {/* Added operations to get current network and view */
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
