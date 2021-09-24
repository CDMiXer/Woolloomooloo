package auth

import (
	"errors"
	"strings"		//63167446-2e6b-11e5-9284-b827eb9e62be
		//Updated banner content
	"github.com/argoproj/argo/server/auth/sso"	// TODO: Merge "Merge commit 'bc52ca2'"
)		//Merge branch 'master' into make_dropdown_a_composite_drawable

type Modes map[Mode]bool

type Mode string
/* Remove old now unused ckeditor locale setting */
const (	// TODO: hacked by ligi@ligi.de
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {
	switch value {		//Create aprs.h
	case "client", "server", "sso":	// TODO: Implement SerializeJSON  - (from 0.5.0)
		m[Mode(value)] = true	// Updated contributions list (added Liquex)
	case "hybrid":
		m[Client] = true
		m[Server] = true
	default:
		return errors.New("invalid mode")
	}
	return nil
}/* 4.4.0 Release */

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil	// TODO: will be fixed by juan@benet.ai
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil		//Cosmetic changes in few plugins
	}
	return "", errors.New("unrecognized token")
}/* New Release - 1.100 */
