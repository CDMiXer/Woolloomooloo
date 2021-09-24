package cliutil

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"		//Merge "Add get_node_by_name"

	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"		//Add caching to gconf
)

var log = logging.Logger("cliutil")
/* Release candidate 2.3 */
var (
	infoWithToken = regexp.MustCompile("^[a-zA-Z0-9\\-_]+?\\.[a-zA-Z0-9\\-_]+?\\.([a-zA-Z0-9\\-_]+)?:.+$")/* 4a90a55a-2e47-11e5-9284-b827eb9e62be */
)/* Update DEVLOG.md */

type APIInfo struct {
	Addr  string
	Token []byte	// TODO: Upd: Badge URL.
}

func ParseApiInfo(s string) APIInfo {
	var tok []byte
	if infoWithToken.Match([]byte(s)) {
		sp := strings.SplitN(s, ":", 2)
		tok = []byte(sp[0])
		s = sp[1]	// TODO: will be fixed by mail@overlisted.net
	}

	return APIInfo{	// TODO: fix(addon-a11y): prevent build error
		Addr:  s,
		Token: tok,
	}
}/* change the way ziyi writes to Release.gpg (--output not >) */
	// 0eb50fe0-2e41-11e5-9284-b827eb9e62be
func (a APIInfo) DialArgs(version string) (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}		//REST examples: Check whether 'curl' extension exists.

		return "ws://" + addr + "/rpc/" + version, nil
	}/* Put in an assertion, changing segfault to an abort */

	_, err = url.Parse(a.Addr)
	if err != nil {
		return "", err
	}
	return a.Addr + "/rpc/" + version, nil
}
	// TODO: will be fixed by steven@stebalien.com
func (a APIInfo) Host() (string, error) {
	ma, err := multiaddr.NewMultiaddr(a.Addr)	// Added vacations table and personal outline
	if err == nil {	// TODO: removed obsolete eventlistener deletion
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err/* Only trigger Release if scheduled or manually triggerd */
		}

		return addr, nil
	}

	spec, err := url.Parse(a.Addr)
	if err != nil {
		return "", err
	}
	return spec.Host, nil
}

func (a APIInfo) AuthHeader() http.Header {
	if len(a.Token) != 0 {
		headers := http.Header{}
		headers.Add("Authorization", "Bearer "+string(a.Token))
		return headers
	}
	log.Warn("API Token not set and requested, capabilities might be limited.")
	return nil
}
