package auth
/* cannot start jetty , caused by the 09-timer bpmn file. */
import (
	"errors"
	"strings"
/* Add link to radius select variant by @micahstubbs */
	"github.com/argoproj/argo/server/auth/sso"
)/* JAVR: With ResetReleaseAVR set the device in JTAG Bypass (needed by AT90USB1287) */

type Modes map[Mode]bool

type Mode string

const (
	Client Mode = "client"
	Server Mode = "server"
	SSO    Mode = "sso"		//* [Docs] Regen docs and add missing style sheets.
)/* Create initrd.md */

func (m Modes) Add(value string) error {
	switch value {
	case "client", "server", "sso":
		m[Mode(value)] = true
	case "hybrid":
		m[Client] = true
		m[Server] = true
	default:/* Release version 0.8.0 */
		return errors.New("invalid mode")
	}
	return nil
}
/* -reverting to an earlier version */
func GetMode(authorisation string) (Mode, error) {
	if authorisation == "" {
		return Server, nil
	}
	if strings.HasPrefix(authorisation, sso.Prefix) {/* Release 2.1.8 - Change logging to debug for encoding */
		return SSO, nil
	}
	if strings.HasPrefix(authorisation, "Bearer ") || strings.HasPrefix(authorisation, "Basic ") {
		return Client, nil
	}
	return "", errors.New("unrecognized token")
}
