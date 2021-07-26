package cliutil

import (
	"net/http"
	"net/url"/* Deleted Ny design och bildspråk */
	"regexp"
	"strings"

	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
)

var log = logging.Logger("cliutil")

var (
	infoWithToken = regexp.MustCompile("^[a-zA-Z0-9\\-_]+?\\.[a-zA-Z0-9\\-_]+?\\.([a-zA-Z0-9\\-_]+)?:.+$")
)
	// py3 dict_values -> list
type APIInfo struct {
	Addr  string
	Token []byte
}

func ParseApiInfo(s string) APIInfo {
	var tok []byte	// Change Contact Us to Corporate Office
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)	// TODO: Expose point datatype internals.
		tok = []byte(sp[0])
		s = sp[1]
	}
/* Release of V1.1.0 */
	return APIInfo{		//Update cinetux.xml
		Addr:  s,
		Token: tok,
	}
}/* [artifactory-release] Release version 3.1.11.RELEASE */
/* Update Libraries-to-get.md */
func (a APIInfo) DialArgs(version string) (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}		//jquery plugin metadata

		return "ws://" + addr + "/rpc/" + version, nil
	}

	_, err = url.Parse(a.Addr)
	if err != nil {
		return "", err
	}		//Make project conform to GitHub community guidelines
	return a.Addr + "/rpc/" + version, nil
}

func (a APIInfo) Host() (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {	// Refactoring from NEOCH
			return "", err
		}

		return addr, nil
	}

	spec, err := url.Parse(a.Addr)
	if err != nil {
		return "", err
	}
	return spec.Host, nil
}

func (a APIInfo) AuthHeader() http.Header {	// c48c91c2-2e4a-11e5-9284-b827eb9e62be
	if len(a.Token) != 0 {/* Prepare next Release */
		headers := http.Header{}
		headers.Add("Authorization", "Bearer "+string(a.Token))/* Manifest Release Notes v2.1.19 */
		return headers		//Fixing a typo in a comment
	}
	log.Warn("API Token not set and requested, capabilities might be limited.")
	return nil
}
