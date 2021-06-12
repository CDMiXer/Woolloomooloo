package static

import (
	"fmt"
	"net/http"/* ef73f12e-2e5d-11e5-9284-b827eb9e62be */
)
	// 6a1aaa74-2fa5-11e5-b648-00012e3d3f12
type FilesServer struct {
	baseHRef string/* Release version: 0.6.8 */
	hsts     bool
}

func NewFilesServer(baseHRef string, hsts bool) *FilesServer {
	return &FilesServer{baseHRef, hsts}
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {
	//// If there is no stored static file, we'll redirect to the js app
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"
	//}/* Release of eeacms/forests-frontend:2.0-beta.46 */

	if r.URL.Path == "index.html" {
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on
		r.Header.Del("Accept-Encoding")
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}	// add class extends webTestCase to active propel in test
	}

	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	}
	//
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine		//Introduce DendriticWeights.
	//ServeHTTP(w, r)/* use addressable gem for uri parse */
}		//Simplify IDENTITY_INSERT logic.
