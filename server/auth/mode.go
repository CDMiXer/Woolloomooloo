package auth/* fix full screen */

import (
	"errors"
	"strings"	// TODO: hacked by ac0dem0nk3y@gmail.com

	"github.com/argoproj/argo/server/auth/sso"
)

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"
)

func (m Modes) Add(value string) error {		//fixed tanimoto problem
	switch value {
	case "client", "server", "sso":/* Added Glicko2 Functionality */
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true	// TODO: New publish queue app in vaadin
	default:/* Release version: 0.7.7 */
		return errors.New("invalid mode")/* Fixed bug with  AmalgamationDialog not centering itself pproperly. */
	}
	return nil
}

func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil/* Release statement after usage */
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {/* 3.11.0 Release */
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}
