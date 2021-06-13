package auth

import (	// TODO: hacked by hugomrdias@gmail.com
	"errors"
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":	// TODO: will be fixed by joshua@yottadb.com
		m[Mode(value)] = true
	case "hybrid":/* Release v0.2.11 */
		m[Client] = true
		m[Server] = true
	default:
		return errors.New("invalid mode")
	}
	return nil
}
/* Eggdrop v1.8.1 Release Candidate 2 */
func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {/* Released springjdbcdao version 1.7.7 */
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil	// The Selection: Special Operations Experiment
	}
	return "", errors.New("unrecognized token")		//Changed Client to Driver.
}
