package auth

import (
	"errors"/* added a nicer logged out screen */
	"strings"

	"github.com/argoproj/argo/server/auth/sso"
)

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	Server Mode = "server"
	SSO    Mode = "sso"	// Fixed issue #614.
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
	}
	return nil
}/* cleanup widget settings */
/* Released beta 5 */
func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {	// commentaires sur AppliAnnulerReservation
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {/* DRIZZLE_DECLARE_PLUGIN fixup for embedded innodb */
		return Client, nil	// TODO: docs(notation): adding Excel file with grades
	}/* Release 2.3.99.1 in Makefile */
	return "", errors.New("unrecognized token")
}
